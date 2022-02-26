package service

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	pb "example.com/bap/blogprofile"
	"example.com/bap/drive"
	"example.com/bap/mongodb"
	"example.com/bap/util/data"
)

type Service interface {
	BlogToBlogDetail(blog *data.Blog) *pb.BlogDetail
	BlogsToBlogList(blogs []data.Blog) *pb.BlogList
	ProfileToProfileDetail(profile *data.Profile) *pb.ProfileDetail

	BlogOpenFilter(blog *data.Blog) *data.Blog
	BlogsOpenFilter(blogs []data.Blog) []data.Blog
	BlogsArticleMask(blogs []data.Blog) []data.Blog
	ZipToHtml(r *http.Response) (string, error)

	ConstructDB(database string) error
	DestructDB() error
	ConstructDrive() error

	Profile() (*data.Profile, error)
	Blog(id string) (*data.Blog, error)
	BlogMeta(id string) (*data.Blog, error)
	Blogs() ([]data.Blog, error)
	NewBlog(blog *data.Blog) error
	UpdateBlog(blog *data.Blog, id string) error
}

type ServiceImpl struct {
	BlogDB       mongodb.DB
	ProfileDB    mongodb.DB
	GoogleDrive  drive.Drive
	ServiceIndex Service
}

type BlogsWrapper struct {
	Total int
	Items []data.Blog
	Tags  []string
}

func NewService() Service {
	return &ServiceImpl{
		BlogDB:       mongodb.NewDB(),
		ProfileDB:    mongodb.NewDB(),
		GoogleDrive:  drive.NewDrive(),
		ServiceIndex: &ServiceImpl{},
	}
}

func (srv *ServiceImpl) BlogToBlogDetail(blog *data.Blog) *pb.BlogDetail {
	return &pb.BlogDetail{
		Id:      blog.Id.Hex(),
		Article: blog.Article,
		Open:    blog.Open,
		Tag:     blog.Tag,
		Title:   blog.Title,
		Date:    blog.Date,
	}
}

func (srv *ServiceImpl) BlogsToBlogList(blogs []data.Blog) *pb.BlogList {
	total := len(blogs)
	blogList := make([]*pb.BlogDetail, total)
	tagDict := make(map[string]int)
	for i, blog := range blogs {
		blogList[i] = srv.BlogToBlogDetail(&blog)
		for _, tag := range blog.Tag {
			tagDict[tag] = 0
		}
	}
	tags := make([]string, len(tagDict))
	i := 0
	for tag := range tagDict {
		tags[i] = tag
		i++
	}

	return &pb.BlogList{
		Total: int32(total),
		Blogs: blogList,
		Tags:  tags,
	}
}

func (srv *ServiceImpl) ProfileToProfileDetail(profile *data.Profile) *pb.ProfileDetail {
	return &pb.ProfileDetail{
		Content: profile.Content,
		Date:    profile.Date,
	}
}

func (srv *ServiceImpl) ConstructDB(database string) error {
	var isErr = false
	// Connect DB
	err := srv.BlogDB.ConstructDB(database, "blog")
	if err != nil {
		isErr = true
	}
	err = srv.ProfileDB.ConstructDB(database, "profile")
	if err != nil {
		isErr = true
	}

	if isErr {
		return err
	} else {
		return nil
	}
}

func (srv *ServiceImpl) DestructDB() error {
	var isErr = false
	// Disconnect DB
	err := srv.BlogDB.DestructDB()
	if err != nil {
		isErr = true
	}
	err = srv.ProfileDB.DestructDB()
	if err != nil {
		isErr = true
	}

	if isErr {
		return err
	} else {
		return nil
	}
}

func (srv *ServiceImpl) ConstructDrive() error {
	var ctx = context.Background()
	err := srv.GoogleDrive.Construct(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (srv *ServiceImpl) BlogOpenFilter(blog *data.Blog) *data.Blog {
	if blog.Open {
		return blog
	} else {
		return nil
	}
}

func (srv *ServiceImpl) BlogsOpenFilter(blogs []data.Blog) []data.Blog {
	var result []data.Blog
	for _, blog := range blogs {
		if srv.BlogOpenFilter(&blog) == nil {
			continue
		}
		result = append(result, blog)
	}
	return result
}

func (srv *ServiceImpl) BlogsArticleMask(blogs []data.Blog) []data.Blog {
	for i := 0; i < len(blogs); i++ {
		blogs[i].Article = ""
	}
	return blogs
}

func (srv *ServiceImpl) ZipToHtml(r *http.Response) (string, error) {
	type imgInfo struct {
		img string
		exp string
	}

	var (
		images  = make(map[string]imgInfo)
		content string
	)
	buff, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(buff)
	zips, err := zip.NewReader(reader, reader.Size())
	if err != nil {
		return "", err
	}
	// extract
	for _, f := range zips.File {
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		b, err := ioutil.ReadAll(rc)
		if err != nil {
			return "", err
		}
		if strings.Contains(f.Name, "html") {
			content = string(b)
		} else {
			encodeStr := base64.StdEncoding.EncodeToString(b)
			splitName := strings.Split(f.Name, ".")
			images[f.Name] = imgInfo{
				img: encodeStr,
				exp: splitName[len(splitName)-1],
			}
		}
	}
	// encode img to src: <img src=...>
	for name, info := range images {
		content = strings.Replace(
			content, name, fmt.Sprintf("data:image/%s;base64,%s", info.exp, info.img), -1,
		)
	}
	return content, nil
}

func (srv *ServiceImpl) Profile() (*data.Profile, error) {
	// Read profile from db
	profile, err := srv.ProfileDB.ReadProfile()
	if err != nil {
		return nil, err
	}
	// Get contents with Google Drive API
	resp, err := srv.GoogleDrive.Export(profile.Content, "application/zip", context.Background())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Zip to HTML
	content, err := srv.ServiceIndex.ZipToHtml(resp)
	if err != nil {
		return nil, err
	}
	profile.Content = content

	return profile, nil
}

func (srv *ServiceImpl) Blog(id string) (*data.Blog, error) {
	// Read blog from db
	blog, err := srv.BlogDB.ReadBlog(id)
	if err != nil {
		return nil, err
	}
	// Get contents with Google Drive API
	resp, err := srv.GoogleDrive.Export(blog.Article, "application/zip", context.Background())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Zip to HTML
	article, err := srv.ServiceIndex.ZipToHtml(resp)
	if err != nil {
		return nil, err
	}
	blog.Article = article

	return blog, nil
}

func (srv *ServiceImpl) BlogMeta(id string) (*data.Blog, error) {
	return srv.BlogDB.ReadBlog(id)
}

func (srv *ServiceImpl) Blogs() ([]data.Blog, error) {
	blogs, err := srv.BlogDB.ReadBlogs()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (srv *ServiceImpl) NewBlog(blog *data.Blog) error {
	return srv.BlogDB.InsertBlog(blog)
}

func (srv *ServiceImpl) UpdateBlog(blog *data.Blog, id string) error {
	return srv.BlogDB.UpdateBlog(blog, id)
}
