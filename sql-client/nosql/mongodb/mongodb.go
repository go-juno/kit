package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/xerrors"
)

type MongoDBConfig struct {
	ApplyURI   string //mongodb://localhost:27017
	DataBase   string //库
	Collection string // 表
}
type MongoDBClient interface {
	InsertOne(ctx context.Context, document interface{}) (string, error)
	FindOneAndDelete(ctx context.Context, filter interface{}) (bool, error)
	DeleteOne(ctx context.Context, document interface{}) (bool, error)
	InsertMany(ctx context.Context, documents []interface{}) ([]string, error)
	DeleteMany(ctx context.Context, documentBson FindBsonM) (int64, error)
	FindOneAndReplace(ctx context.Context, document interface{}, replaceDocument interface{}) (bool, error)
	UpdateMany(ctx context.Context, filter FindBsonM, document FindBsonM) (int64, error)
	FindOne(ctx context.Context, filterModel interface{}, setSkip int64, outModel interface{}) (interface{}, error)
	Find(ctx context.Context, FilterBson interface{}, findOptions FindOptions, outModel []interface{}) ([]interface{}, error)
	Aggregate(ctx context.Context, pipelineBson FindBson, outModel []interface{}) ([]interface{}, error)
}

type FindBson struct {
	BsonD bson.D
}
type FindBsonM struct {
	BsonM bson.M
}

type FindBsonE struct {
	BsonE bson.E
}

type FindBsonA struct {
	BsonA bson.A
}

type FindBsonRaw struct {
	BsonRaw bson.Raw
}

type FindOptions struct {
	setSKip  int64
	setLimit int64
}

type monGODBClient struct {
	collection *mongo.Collection
}

func (m *monGODBClient) InsertOne(ctx context.Context, document interface{}) (string, error) {
	one, err := m.collection.InsertOne(ctx, document)
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return one.InsertedID.(primitive.ObjectID).Hex(), nil

}

func (m *monGODBClient) Aggregate(ctx context.Context, pipelineBson FindBson, outModel []interface{}) ([]interface{}, error) {
	cursor, err := m.collection.Aggregate(ctx, mongo.Pipeline{pipelineBson.BsonD})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			return
		}
	}()
	if err = cursor.All(context.TODO(), &outModel); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return outModel, nil
}

func (m *monGODBClient) Find(ctx context.Context, FilterBson interface{}, findOptions FindOptions, outModel []interface{}) ([]interface{}, error) {
	cursor, err := m.collection.Find(ctx, FilterBson, options.Find().SetSkip(findOptions.setSKip).SetLimit(findOptions.setLimit))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			return
		}
	}()
	if err = cursor.All(context.TODO(), &outModel); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return outModel, nil

}
func (m *monGODBClient) FindOne(ctx context.Context, filterModel interface{}, setSkip int64, outModel interface{}) (interface{}, error) {
	err := m.collection.FindOne(ctx, filterModel).Decode(outModel)
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return outModel, nil

}
func (m *monGODBClient) UpdateMany(ctx context.Context, filter FindBsonM, document FindBsonM) (int64, error) {
	update, err := m.collection.UpdateMany(ctx, filter, document)
	if err != nil {
		return -1, xerrors.Errorf("%w", err)
	}
	return update.MatchedCount, nil

}
func (m *monGODBClient) FindOneAndReplace(ctx context.Context, document interface{}, replaceDocument interface{}) (bool, error) {
	err := m.collection.FindOneAndReplace(ctx, document, replaceDocument).Err()
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	return true, nil

}

func (m *monGODBClient) DeleteMany(ctx context.Context, documentBson FindBsonM) (int64, error) {
	result, err := m.collection.DeleteMany(ctx, documentBson.BsonM)
	if err != nil {
		return -1, xerrors.Errorf("%w", err)
	}

	return result.DeletedCount, nil

}

func (m *monGODBClient) InsertMany(ctx context.Context, documents []interface{}) ([]string, error) {
	var objectIDs []string
	many, err := m.collection.InsertMany(ctx, documents)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	for _, insertedID := range many.InsertedIDs {
		id := insertedID.(primitive.ObjectID)
		objectIDs = append(objectIDs, id.Hex())
	}
	return objectIDs, nil

}
func (m *monGODBClient) DeleteOne(ctx context.Context, document interface{}) (bool, error) {
	_, err := m.collection.DeleteOne(ctx, document)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	return true, nil

}

func (m *monGODBClient) FindOneAndDelete(ctx context.Context, filter interface{}) (bool, error) {
	err := m.collection.FindOneAndDelete(ctx, filter).Err()
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	return true, err

}

func NewMongoDBClient(config MongoDBConfig) (MongoDBClient, error) {

	//1.建立连接
	dClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.ApplyURI).SetConnectTimeout(5*time.Second))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	ctx, canFn := context.WithTimeout(context.Background(), time.Second*10)
	defer canFn()
	if err = dClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	//2.选择数据库 your_db
	db := dClient.Database(config.DataBase)
	//3.选择表 your_collection
	collection := db.Collection(config.Collection)
	return &monGODBClient{
		collection: collection,
	}, nil
}
