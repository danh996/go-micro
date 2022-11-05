package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"

	"github.com/danh996/go-micro/logger-service/data"
	"github.com/danh996/go-micro/logger-service/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	//write the log

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.InsertOne(logEntry)

	if err != nil {
		res := &logs.LogResponse{
			Result: "fail",
		}
		return res, err
	}

	// return response
	res := &logs.LogResponse{
		Result: "logged",
	}
	return res, err

}

func (app *Config) grpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))

	if err != nil {
		log.Fatalf("failed to listen to grpc: %v", err)
	}

	s := grpc.NewServer()
	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Printf("Grpc Server started on port %s", gRpcPort)

	if err := s.Server(lis); err != nil {
		log.Fatalf("failed to listen to grpc: %v", err)
	}
}
