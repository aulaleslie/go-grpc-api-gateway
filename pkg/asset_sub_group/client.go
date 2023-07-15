package asset_sub_group

import (
	"fmt"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_sub_group/pb"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AssetSubGroupServiceClient
}

func InitServiceClient(c *config.Config) pb.AssetSubGroupServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AssetSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAssetSubGroupServiceClient(cc)
}
