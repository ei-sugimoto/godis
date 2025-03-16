package server

import (
	"context"
	"log"
	"net"

	proto "github.com/ei-sugimoto/godis/internal/gen/go/proto/v1"
	"github.com/ei-sugimoto/godis/internal/pkg/config"
	"github.com/ei-sugimoto/godis/internal/pkg/env"
	"github.com/ei-sugimoto/godis/internal/pkg/err"
	"github.com/ei-sugimoto/godis/internal/pkg/server/middleware"
	"github.com/ei-sugimoto/godis/internal/pkg/service"
	"github.com/ei-sugimoto/godis/internal/pkg/store"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GodisServe struct {
	ln net.Listener
}

func NewGodisServe() *GodisServe {
	return &GodisServe{}
}

func (g *GodisServe) Listen() error {
	port, err := env.GetPort()
	if err != nil {
		return err
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	g.ln = ln
	log.Println("Listening on port:", port)
	return nil
}

func (g *GodisServe) Serve() error {
	if g.ln == nil {
		return err.ErrNoListener
	}

	pt, err := env.GetConfigPath()
	if err != nil {
		return err
	}

	c, err := config.ParseConfig(pt)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,
			middleware.KeyInterceptor(c.APIKey),
		),
	)

	db := store.NewDB()
	proto.RegisterRecordServiceServer(s, service.NewRecordService(db))

	proto.RegisterPingServiceServer(s, service.NewPingService())

	reflection.Register(s)

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		log.Println("Starting server...")
		defer log.Println("Stopping server...")
		return s.Serve(g.ln)
	})

	eg.Go(func() error {
		<-ctx.Done()
		s.GracefulStop()
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
