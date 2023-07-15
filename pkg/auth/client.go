package auth

import (
	"fmt"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
