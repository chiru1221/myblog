package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"example.com/bap/util/data"
	"example.com/bap/util/secrets"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB interface {
	UpdateBlog(blog *data.Blog, id string) error
	InsertBlog(blog *data.Blog) error
	ReadBlogs() ([]data.Blog, error)
	ReadBlog(id string) (*data.Blog, error)
	ReadProfile() (*data.Profile, error)

	ConstructDB(database, collection string) error
	DestructDB() error
}

type DBImpl struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewDB() DB {
	return &DBImpl{
		Client:     nil,
		Collection: nil,
	}
}

func toJapanTime(t time.Time) string {
	var loc *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)
	return t.In(loc).Format("2006/01/02")
}

func (db *DBImpl) ConstructDB(database, collection string) error {
	var (
		mongoUser   string = secrets.SecretsFile(os.Getenv("MONGO_INITDB_ROOT_USERNAME_FILE"))
		mongoPasswd string = secrets.SecretsFile(os.Getenv("MONGO_INITDB_ROOT_PASSWORD_FILE"))
	)
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@db", mongoUser, mongoPasswd)))
	if err != nil {
		return err
	}
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db.Client = client
	db.Collection = client.Database(database).Collection(collection)
	return nil
}

func (db *DBImpl) DestructDB() error {
	// Disconnect DB
	return db.Client.Disconnect(context.TODO())
}

func (db *DBImpl) UpdateBlog(blog *data.Blog, id string) error {
	if !primitive.IsValidObjectID(id) {
		return fmt.Errorf("Invalid id")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateItem := bson.D{
		{"article", blog.Article},
		{"open", blog.Open},
		{"tag", blog.Tag},
		{"title", blog.Title},
	}

	err = db.Collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": objectId},
		bson.D{
			{"$set", updateItem},
		},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&blog)

	return err
}

func (db *DBImpl) InsertBlog(blog *data.Blog) error {
	insertItem := bson.D{
		{"article", blog.Article},
		{"open", blog.Open},
		{"tag", blog.Tag},
		{"title", blog.Title},
	}

	insertedResult, err := db.Collection.InsertOne(context.Background(), insertItem)
	if err != nil {
		return err
	}
	blog.Id = insertedResult.InsertedID.(primitive.ObjectID)

	return nil
}

func (db *DBImpl) ReadBlogs() ([]data.Blog, error) {
	var blogs []data.Blog

	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		doc := data.Blog{}
		cursor.Decode(&doc)
		blogs = append(blogs, doc)
	}
	return blogs, nil
}

func (db *DBImpl) ReadBlog(id string) (*data.Blog, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, fmt.Errorf("Invalid id")
	}
	var blog *data.Blog
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := db.Collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objectID}}).Decode(&blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (db *DBImpl) ReadProfile() (*data.Profile, error) {
	var profile *data.Profile
	err := db.Collection.FindOne(context.Background(), bson.M{}).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
