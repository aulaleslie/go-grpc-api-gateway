package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_sub_group/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.AssetSubGroupServiceClient) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: int32(id),
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
