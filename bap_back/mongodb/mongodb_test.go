package mongodb_test

import (
	"testing"

	"example.com/bap/mongodb"
	"example.com/bap/util/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUpdateBlog(t *testing.T) {
	type testCase struct {
		name       string
		id         primitive.ObjectID
		input      *data.Blog
		wantResult data.Blog
	}

	var (
		db mongodb.DBImpl
		mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	)
	defer mt.Close()

	objectID := primitive.NewObjectID()
	tests := []testCase{
		{
			name: "All value",
			id:   objectID,
			input: &data.Blog{
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    false,
				Date:    "2022/01/01",
			},
			wantResult: data.Blog{
				Id:      objectID,
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    false,
				Date:    "2022/01/01",
			},
		},
		{
			name:  "No value",
			id:    objectID,
			input: &data.Blog{},
			wantResult: data.Blog{
				Id: objectID,
			},
		},
	}

	for _, tc := range tests {
		mt.Run(tc.name, func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bson.D{
					{"_id", tc.id},
					{"article", tc.input.Article},
					{"open", tc.input.Open},
					{"tag", tc.input.Tag},
					{"title", tc.input.Title},
					{"date", tc.input.Date},
				}},
			})

			db = mongodb.DBImpl{Collection: mt.Coll}

			err := db.UpdateBlog(tc.input, tc.id.Hex())
			if err != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, nil, err)
			}

			if tc.input.Title != tc.wantResult.Title {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Title, tc.input.Title)
			}
			if len(tc.input.Tag) != len(tc.wantResult.Tag) {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, tc.input.Tag)
			} else {
				for i := 0; i < len(tc.input.Tag); i++ {
					if tc.input.Tag[i] != tc.wantResult.Tag[i] {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, tc.input.Tag)
					}
				}
			}
			if tc.input.Article != tc.wantResult.Article {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Article, tc.input.Article)
			}
			if tc.input.Open != tc.wantResult.Open {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Open, tc.input.Open)
			}
			if tc.input.Date != tc.wantResult.Date {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Date, tc.input.Date)
			}

			mt.ResetClient(nil)
		})
	}

}

func TestInsertBlog(t *testing.T) {
	type testCase struct {
		name       string
		input      *data.Blog
		wantResult data.Blog
	}

	var (
		db mongodb.DBImpl
		mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	)
	defer mt.Close()

	tests := []testCase{
		{
			name: "All value",
			input: &data.Blog{
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    false,
				Date:    "2022/01/01",
			},
			wantResult: data.Blog{
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    false,
				Date:    "2022/01/01",
			},
		},
		{
			name:       "No value",
			input:      &data.Blog{},
			wantResult: data.Blog{},
		},
	}

	for _, tc := range tests {
		mt.Run(tc.name, func(mt *mtest.T) {
			mt.AddMockResponses(mtest.CreateSuccessResponse())

			db = mongodb.DBImpl{Collection: mt.Coll}
			err := db.InsertBlog(tc.input)

			if err != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, nil, err)
			}

			if tc.input.Title != tc.wantResult.Title {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Title, tc.input.Title)
			}
			if len(tc.input.Tag) != len(tc.wantResult.Tag) {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, tc.input.Tag)
			} else {
				for i := 0; i < len(tc.input.Tag); i++ {
					if tc.input.Tag[i] != tc.wantResult.Tag[i] {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, tc.input.Tag)
					}
				}
			}
			if tc.input.Article != tc.wantResult.Article {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Article, tc.input.Article)
			}
			if tc.input.Open != tc.wantResult.Open {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Open, tc.input.Open)
			}
			if tc.input.Date != tc.wantResult.Date {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Date, tc.input.Date)
			}

			mt.ResetClient(nil)
		})
	}
}

func TestReadBlogs(t *testing.T) {
	type testCase struct {
		name       string
		input      []data.Blog
		wantResult []data.Blog
	}

	var (
		db mongodb.DBImpl
		mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	)
	defer mt.Close()

	tests := []testCase{
		{
			name: "Single data",
			input: []data.Blog{
				{
					Id:      primitive.NewObjectID(),
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    true,
					Date:    "2022/01/01",
				},
			},
			wantResult: []data.Blog{
				{
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    true,
					Date:    "2022/01/01",
				},
			},
		},
		{
			name: "Multiple data",
			input: []data.Blog{
				{
					Id:      primitive.NewObjectID(),
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    true,
					Date:    "2022/01/01",
				},
				{
					Id:      primitive.NewObjectID(),
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    false,
					Date:    "2022/01/01",
				},
			},
			wantResult: []data.Blog{
				{
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    true,
					Date:    "2022/01/01",
				},
				{
					Id:      primitive.NewObjectID(),
					Title:   "hoge",
					Tag:     []string{"hoge"},
					Article: "hoge",
					Open:    false,
					Date:    "2022/01/01",
				},
			},
		},
	}

	for _, tc := range tests {
		mt.Run(tc.name, func(mt *mtest.T) {
			var batch mtest.BatchIdentifier
			for i, input_i := range tc.input {
				if i == 0 {
					batch = mtest.FirstBatch
				} else {
					batch = mtest.NextBatch
				}
				mt.AddMockResponses(mtest.CreateCursorResponse(int64(i+1), "test.hoge", batch, bson.D{
					{"_id", input_i.Id},
					{"title", input_i.Title},
					{"tag", input_i.Tag},
					{"article", input_i.Article},
					{"open", input_i.Open},
					{"date", input_i.Date},
				}))
			}

			db = mongodb.DBImpl{Collection: mt.Coll}
			blogs, err := db.ReadBlogs()
			if err != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, nil, err)
			}

			if len(blogs) != len(tc.wantResult) {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, blogs)
			} else {
				for i, wantResult_i := range tc.wantResult {
					if blogs[i].Title != wantResult_i.Title {
						t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, wantResult_i.Title, blogs[i].Title)
					}
					if len(blogs[i].Tag) != len(wantResult_i.Tag) {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, wantResult_i.Tag, blogs[i].Tag)
					} else {
						for j := 0; j < len(blogs[i].Tag); j++ {
							if blogs[i].Tag[j] != wantResult_i.Tag[j] {
								t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, wantResult_i.Tag, blogs[i].Tag)
							}
						}
					}
					if blogs[i].Article != wantResult_i.Article {
						t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, wantResult_i.Article, blogs[i].Article)
					}
					if blogs[i].Open != wantResult_i.Open {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, wantResult_i.Open, blogs[i].Open)
					}
					if blogs[i].Date != wantResult_i.Date {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, wantResult_i.Date, blogs[i].Date)
					}
				}
			}

			mt.ResetClient(nil)
		})
	}
}

func TestReadBlog(t *testing.T) {
	type testCase struct {
		name       string
		input      data.Blog
		wantResult data.Blog
	}

	var (
		db mongodb.DBImpl
		mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	)
	defer mt.Close()

	tests := []testCase{
		{
			name: "Exist",
			input: data.Blog{
				Id:      primitive.NewObjectID(),
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    true,
				Date:    "2022/01/01",
			},
			wantResult: data.Blog{
				Title:   "hoge",
				Tag:     []string{"hoge"},
				Article: "hoge",
				Open:    true,
				Date:    "2022/01/01",
			},
		},
		{
			name:       "No exist",
			input:      data.Blog{},
			wantResult: data.Blog{},
		},
	}

	for _, tc := range tests {
		mt.Run(tc.name, func(mt *mtest.T) {
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.hoge", mtest.FirstBatch, bson.D{
				{"_id", tc.input.Id},
				{"title", tc.input.Title},
				{"tag", tc.input.Tag},
				{"article", tc.input.Article},
				{"open", tc.input.Open},
				{"date", tc.input.Date},
			}))

			db = mongodb.DBImpl{Collection: mt.Coll}
			blog, err := db.ReadBlog(tc.input.Id.Hex())
			if err != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, nil, err)
				return
			}
			if blog.Title != tc.wantResult.Title {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Title, blog.Title)
			}
			if len(blog.Tag) != len(tc.wantResult.Tag) {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, blog.Tag)
			} else {
				for i := 0; i < len(blog.Tag); i++ {
					if blog.Tag[i] != tc.wantResult.Tag[i] {
						t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, blog.Tag)
					}
				}
			}
			if blog.Article != tc.wantResult.Article {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Article, blog.Article)
			}
			if blog.Open != tc.wantResult.Open {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Open, blog.Open)
			}
			if blog.Date != tc.wantResult.Date {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Date, blog.Date)
			}

			mt.ResetClient(nil)
		})
	}
}

func TestReadProfile(t *testing.T) {
	type testCase struct {
		name       string
		input      data.Profile
		wantResult data.Profile
	}

	var (
		db mongodb.DBImpl
		mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	)
	defer mt.Close()

	tests := []testCase{
		{
			name:       "Exist",
			input:      data.Profile{Content: "hoge", Date: "2022/01/01"},
			wantResult: data.Profile{Content: "hoge", Date: "2022/01/01"},
		},
		{
			name:       "No exist",
			input:      data.Profile{},
			wantResult: data.Profile{},
		},
	}

	for _, tc := range tests {
		mt.Run(tc.name, func(mt *mtest.T) {
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.hoge", mtest.FirstBatch, bson.D{
				{"content", tc.input.Content},
				{"date", tc.input.Date},
			}))

			db = mongodb.DBImpl{Collection: mt.Coll}
			profile, err := db.ReadProfile()
			if err != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, nil, err)
			}

			if profile.Content != tc.wantResult.Content {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Content, profile.Content)
			}
			if profile.Date != tc.wantResult.Date {
				t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Date, profile.Date)
			}
			mt.ResetClient(nil)
		})
	}
}
