package main

import (
	"context"
	user "github.com/0xhuk/bazel-go-grpc-example/bazel-bin/grpc-gateway/proto"
	"github.com/0xhuk/bazel-go-grpc-example/bazel-bin/pure-proto/proto/common"
	"google.golang.org/grpc"
	"log"
)
const (
	HttpAddr = ":4501"
	GrpcAddr = ":9091"
)
func main() {
	var err error
	var conn = new(grpc.ClientConn)
	var resp = new(common.ApiResult)
	var message = &user.GetUserMessageListResponse{}

	if conn, err = grpc.Dial(GrpcAddr, grpc.WithInsecure()); err != nil {
		log.Panicf("grpc dial err: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)

	if resp, err = client.GetUserMessageList(context.Background(), &user.GetUserMessageListRequest{}); err != nil {
		log.Panicf("server throw err: %v", err)
	}
	if err = resp.Data.UnmarshalTo(message); err != nil {
		log.Panicf("any unwrap err: %v", err)
	}
	log.Printf("res: %s", message)
}
