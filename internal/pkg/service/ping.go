package service

import (
	"context"

	pingv1 "github.com/ei-sugimoto/godis/internal/gen/go/proto/v1"
)

type PingService struct {
	pingv1.UnimplementedPingServiceServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (p *PingService) Ping(ctx context.Context, req *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return &pingv1.PingResponse{Message: "pong"}, nil
}
