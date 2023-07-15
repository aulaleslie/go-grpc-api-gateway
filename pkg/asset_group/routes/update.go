package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_group/pb"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context, c pb.AssetGroupServiceClient) {
	body := CreateAssetGroupRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb.CreateUpdateRequest{
		Id:               int32(id),
		AgrGroupName:     body.AgrGroupName,
		AgrBusinessGroup: body.AgrBusinessGroup,
		SubGroup:         body.SubGroup,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if !res.Status {
		ctx.AbortWithError(int(res.Data.Code), errors.New(res.Data.Message))
	}

	ctx.JSON(int(res.Data.Code), &res)
}
