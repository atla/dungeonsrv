package db

import (
	"context"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "dungeonsrv"
)

// Client database
type Client struct {
	client    *mongo.Client
	Connected bool
}

// New creates a new Client object
func New() *Client {
	return &Client{
		Connected: false,
	}
}

// IDFromObjectID ...
func (dba *Client) IDFromObjectID(bid interface{}) string {

	return bid.(primitive.ObjectID).Hex()

}

// Close connection
func (dba *Client) Close() error {
	return dba.client.Disconnect(context.TODO())
}

//C opens a given collection in the xdcc database
func (dba *Client) C(coll string) *mongo.Collection {
	return dba.client.Database(database).Collection(coll)
}

//FindAll returns all entities of a given collection
func (dba *Client) FindAll(coll string) (*mongo.Cursor, error) {
	findFilter := bson.D{}
	return dba.C(coll).Find(context.TODO(), findFilter)
}

//Find returns all entities of a given collection
func (dba *Client) Find(coll string, key string, value string) (*mongo.Cursor, error) {
	findFilter := bson.D{{key, value}}

	// findFilter2 := bson.D{{
	// 	key,
	// 	bson.D{{
	// 		"$in",
	// 		bson.E{value}
	// 	}}
	// }}

	return dba.C(coll).Find(context.TODO(), findFilter)
}

//FindOne returns all entities of a given collection
func (dba *Client) FindOne(coll string, key string, value string) *mongo.SingleResult {

	return dba.C(coll).FindOne(context.TODO(), bson.D{{key, value}})
}

//FindByID returns all entities of a given collection
func (dba *Client) FindByID(coll string, id string) *mongo.SingleResult {
	objID, _ := primitive.ObjectIDFromHex(id)
	return dba.C(coll).FindOne(context.TODO(), bson.D{{"_id", objID}})
}

//DeleteByID returns all entities of a given collection
func (dba *Client) DeleteByID(coll string, id string) (*mongo.DeleteResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return dba.C(coll).DeleteOne(context.TODO(), bson.D{{"_id", objID}})
}

//InsertOne inserts one document
func (dba *Client) InsertOne(coll string, data interface{}) (*mongo.InsertOneResult, error) {
	return dba.client.Database(database).Collection(coll).InsertOne(context.TODO(), data)
}

//UpdateOne inserts one document
func (dba *Client) UpdateOne(coll string, key string, value string, data interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{key: value}
	update := bson.M{"$set": data}

	return dba.client.Database(database).Collection(coll).UpdateOne(context.TODO(), filter, update)
}

//UpdateOneByID inserts one document
func (dba *Client) UpdateOneByID(coll string, id primitive.ObjectID, data interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}

	return dba.client.Database(database).Collection(coll).UpdateOne(context.TODO(), filter, update)
}

//Connect Connects to the xdcc database
func (dba *Client) Connect(host string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(host)

	log.WithField("host", host).Info("Connecting to MongoDB instance")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	dba.client = client
	dba.Connected = true
}
