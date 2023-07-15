package asset_group

import (
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_group/routes"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/auth"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/assets/assets-group")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.Create)
	routes.GET("/", svc.Search)
	routes.PUT("/:id", svc.Update)
	routes.DELETE("/:id", svc.Delete)
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) Search(ctx *gin.Context) {
	routes.Search(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}
