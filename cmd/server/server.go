package main

import (
	"context"
	"fmt"
	user "github.com/0xhuk/bazel-go-grpc-example/bazel-bin/grpc-gateway/proto"
	"github.com/0xhuk/bazel-go-grpc-example/bazel-bin/pure-proto/proto/common"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net"
	"net/http"
)

const (
	HttpAddr = ":4501"
	GrpcAddr = ":9091"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u UserService) GetUserMessageList(ctx context.Context, request *user.GetUserMessageListRequest) (*common.ApiResult, error) {
	var err error
	var result = new(common.ApiResult)
	var message = new(user.GetUserMessageListResponse)
	var anyInstance = new(anypb.Any)
	message.Data = append(message.Data, &user.UserMessage{Message: "1111"})
	if anyInstance, err = anypb.New(message); err != nil {
		return nil, err
	}
	result.Data = anyInstance
	log.Printf("return %+v ", result)
	return result, nil
}

func RunGrpcHttpGateway(grpcServerEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Printf("run grpc http proxy server in %s \n", HttpAddr)
	return http.ListenAndServe(HttpAddr, mux)
}

func main() {
	lis, err := net.Listen("tcp", GrpcAddr)

	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &UserService{})

	reflection.Register(grpcServer)

	go func() {
		fmt.Printf("run grpc server in %s \n", GrpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to grpc serve: %v\n", err)
		}
	}()

	if err := RunGrpcHttpGateway(GrpcAddr); err != nil {
		log.Panicf("failed to gateway serve: %v\n", err)
	}
}
