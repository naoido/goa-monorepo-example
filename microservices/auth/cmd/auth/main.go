package main

import (
	"context"
	"goa-example/microservices/auth"
	auther "goa-example/microservices/auth/gen/auth"
	genpb "goa-example/microservices/auth/gen/grpc/auth/pb"
	genserver "goa-example/microservices/auth/gen/grpc/auth/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("リッスンに失敗しました: %v", err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	svc := auth.NewAuthService()
	endpoints := auther.NewEndpoints(svc)

	genpb.RegisterAuthServer(srv, genserver.New(endpoints, nil))

	reflection.Register(srv)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("gRPCサーバーをシャットダウンしています...")
		srv.GracefulStop()
	}()

	log.Printf("gRPCサーバーが:8080でリッスンしています")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("サービスの提供に失敗しました: %v", err)
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("%sを処理中", info.FullMethod)
	return handler(ctx, req)
}
