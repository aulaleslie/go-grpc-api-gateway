package routes

import (
	"context"
	"net/http"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_sub_group/pb"
	"github.com/gin-gonic/gin"
)

type SearchParams struct {
	Expand  []string `form:"expand"`
	Page    int      `form:"page"`
	PerPage int      `form:"per_page"`
	Search  string   `form:"search"`
	Sort    string   `form:"sort"`
}

func Search(ctx *gin.Context, c pb.AssetSubGroupServiceClient) {
	var params SearchParams
	if err := ctx.BindQuery(&params); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	res, err := c.Read(context.Background(), &pb.ReadRequest{
		Expand:  params.Expand,
		Page:    int32(params.Page),
		PerPage: int32(params.PerPage),
		Search:  params.Search,
		Sort:    params.Sort,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
