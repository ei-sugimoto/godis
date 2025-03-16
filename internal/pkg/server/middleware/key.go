package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// KeyInterceptor checks for a valid API key in the request metadata.
func KeyInterceptor(apiKey string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		if len(md["api_key"]) == 0 || md["api_key"][0] != apiKey {
			return nil, status.Errorf(codes.Unauthenticated, "invalid API key")
		}

		return handler(ctx, req)
	}
}
