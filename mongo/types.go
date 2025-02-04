package mongo

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database interface {
	Collection(string) Collection
	Client() Client
}
type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (interface{}, error)
	Find(context.Context, interface{}, ...options.Lister[options.FindOptions]) (Cursor, error)
	UpdateOne(context.Context, interface{}, interface{}, ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)
	Aggregate(context.Context, interface{}) (Cursor, error)
}
type SingleResult interface {
	Decode(interface{}) error
}
type Client interface {
	Database(string) Database
	StartSession() (mongo.Session, error)
	Disconnect(context.Context) error
	UseSession(ctx context.Context, fn func(context.Context) error) error
	Ping(ctx context.Context) error
}
type Cursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	All(context.Context, interface{}) error
}
type mongoClient struct {
	cl *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}
type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	mc *mongo.Cursor
}
type nullawareDecoder struct { //nolint:all
	defDecoder bsoncodec.ValueDecoder
	zeroValue  reflect.Value
}
