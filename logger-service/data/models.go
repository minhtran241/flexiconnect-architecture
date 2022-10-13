package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Data      string             `bson:"data,omitempty" json:"data,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(entry *LogEntry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(ctx, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into logs:", err)
		return nil
	}
	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})
	cur, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cur.Close(ctx)
	var logs []*LogEntry

	for cur.Next(ctx) {
		var result LogEntry
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Error decoding log into slice:", err)
			return nil, err
		} else {
			logs = append(logs, &result)
		}
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return logs, nil
}
