package main

import (
	"context"
	"log"
	"time"

	"github.com/danh996/go-micro/logger-service/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payLoad RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payLoad.Name,
		Data:      payLoad.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Process payload via RPC: " + payLoad.Name

	return nil
}
