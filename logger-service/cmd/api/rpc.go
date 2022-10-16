package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minhtran241/log-service/data"
)

// RPCServer is the type for the RPC Server. Methods that take this as a receiver are available over RPC, as long as they exported
type RPCServer struct {
}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo writes our payload to MongoDB
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(ctx, data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		log.Println("error writing to MongoDB", err)
		return err
	}

	// resp is the message sent back to the RPC caller
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
