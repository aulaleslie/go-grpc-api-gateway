package routes

import (
	"context"
	"net/http"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_group/pb"
	"github.com/gin-gonic/gin"
)

type CreateAssetGroupRequestBody struct {
	AgrGroupName     string   `json:"agr_group_name"`
	AgrBusinessGroup []string `json:"agr_business_group"`
	SubGroup         []string `json:"subGroup"`
}

func Create(ctx *gin.Context, c pb.AssetGroupServiceClient) {
	body := CreateAssetGroupRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Create(context.Background(), &pb.CreateUpdateRequest{
		AgrGroupName:     body.AgrGroupName,
		AgrBusinessGroup: body.AgrBusinessGroup,
		SubGroup:         body.SubGroup,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
