package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_sub_group/pb"
	"github.com/gin-gonic/gin"
)

type CreateAssetSubGroupRequestBody struct {
	AsgName        string   `json:"asg_name"`
	AsgParentGroup []string `json:"asg_parent_group"`
}

func Create(ctx *gin.Context, c pb.AssetSubGroupServiceClient) {
	body := CreateAssetSubGroupRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Create(context.Background(), &pb.CreateUpdateRequest{
		AsgName:        body.AsgName,
		AsgParentGroup: body.AsgParentGroup,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if !res.Status {
		ctx.AbortWithError(int(res.Data.Code), errors.New(res.Data.Message))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
