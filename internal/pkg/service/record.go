package service

import (
	"context"

	recordv1 "github.com/ei-sugimoto/godis/internal/gen/go/proto/v1"
	"github.com/ei-sugimoto/godis/internal/pkg/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecordService struct {
	DB *store.DB
	recordv1.UnimplementedRecordServiceServer
}

func NewRecordService(db *store.DB) *RecordService {
	return &RecordService{DB: db}
}

func (r *RecordService) Get(ctx context.Context, req *recordv1.GetRequest) (*recordv1.GetResponse, error) {
	v, ok := r.DB.Get(req.Key)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "key %s not found", req.Key)
	}
	return &recordv1.GetResponse{Value: v}, nil
}

func (r *RecordService) Set(ctx context.Context, req *recordv1.SetRequest) (*recordv1.SetResponse, error) {
	r.DB.Set(req.Key, req.Value)
	return &recordv1.SetResponse{Ok: true}, nil
}

func (r *RecordService) Bulk(ctx context.Context, req *recordv1.BulkRequest) (*recordv1.BulkResponse, error) {
	for _, kv := range req.Requests {
		r.DB.Set(kv.Key, kv.Value)
	}
	return &recordv1.BulkResponse{Ok: true}, nil
}
