package asset_group

import (
	"fmt"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_group/pb"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AssetGroupServiceClient
}

func InitServiceClient(c *config.Config) pb.AssetGroupServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AssetSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAssetGroupServiceClient(cc)
}
