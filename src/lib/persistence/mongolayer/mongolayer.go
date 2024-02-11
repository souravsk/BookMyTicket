// src/lib/persistence/mongolayer/mongolayer.go

package mongolayer

import (
	"context"

	"github.com/souravsk/BookMyTicket/src/lib/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		client: client,
	}, nil
}

func (mongoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	collection := mongoLayer.client.Database(DB).Collection(EVENTS)

	if !e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	if !e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}

	result, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return nil, err
	}

	return []byte(result.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (mongoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	collection := mongoLayer.client.Database(DB).Collection(EVENTS)

	objectID, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return persistence.Event{}, err
	}

	filter := bson.M{"_id": objectID}

	var event persistence.Event
	err = collection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		return persistence.Event{}, err
	}

	return event, nil
}

func (mongoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	collection := mongoLayer.client.Database(DB).Collection(EVENTS)

	filter := bson.M{"name": name}

	var event persistence.Event
	err := collection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		return persistence.Event{}, err
	}

	return event, nil
}

func (mongoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	collection := mongoLayer.client.Database(DB).Collection(EVENTS)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var events []persistence.Event
	if err := cursor.All(context.Background(), &events); err != nil {
		return nil, err
	}

	return events, nil
}
