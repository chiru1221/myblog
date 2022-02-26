package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "example.com/bap/blogprofile"
	"example.com/bap/service"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBapServer
	Srv service.Service
}

func (server *Server) Blog(ctx context.Context, blogId *pb.BlogId) (*pb.BlogDetail, error) {
	id := blogId.GetId()
	if id == "" {
		log.Println("Get id error")
		return nil, errors.New("Invalid BlogId")
	}
	blog, err := server.Srv.Blog(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	blog = server.Srv.BlogOpenFilter(blog)
	return server.Srv.BlogToBlogDetail(blog), nil
}

func (server *Server) Blogs(ctx context.Context, noId *pb.NoId) (*pb.BlogList, error) {
	blogs, err := server.Srv.Blogs()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	blogs = server.Srv.BlogsOpenFilter(blogs)

	// Mask article
	blogs = server.Srv.BlogsArticleMask(blogs)
	return server.Srv.BlogsToBlogList(blogs), nil
}

func (server *Server) Profile(ctx context.Context, noId *pb.NoId) (*pb.ProfileDetail, error) {
	profile, err := server.Srv.Profile()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return server.Srv.ProfileToProfileDetail(profile), nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := Server{Srv: service.NewService()}
	if err = server.Srv.ConstructDB("database"); err != nil {
		log.Fatal(err)
	}
	if err = server.Srv.ConstructDrive(); err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBapServer(grpcServer, &server)
	log.Println("Start Server")
	grpcServer.Serve(lis)
}
