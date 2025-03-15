package server

import (
	"context"
	"log"
	"net"

	pingv1 "github.com/ei-sugimoto/godis/internal/gen/go/proto/v1"
	"github.com/ei-sugimoto/godis/internal/pkg/env"
	"github.com/ei-sugimoto/godis/internal/pkg/err"
	"github.com/ei-sugimoto/godis/internal/pkg/service"
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

	s := grpc.NewServer()

	pingv1.RegisterPingServiceServer(s, service.NewPingService())

	reflection.Register(s)

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return s.Serve(g.ln)
	})

	eg.Go(func() error {
		<-ctx.Done()
		s.GracefulStop()
		return nil
	})

	return eg.Wait()
}
