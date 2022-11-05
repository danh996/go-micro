package main

import (
	"context"
	"logger-service/data"
	"logger-service/logs"

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
