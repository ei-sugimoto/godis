package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor is a gRPC middleware that logs the details of each request.
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Received request: %v", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error handling request: %v", status.Convert(err).Message())
	} else {
		log.Printf("Successfully handled request")
	}
	return resp, err
}
