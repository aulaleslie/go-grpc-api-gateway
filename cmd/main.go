package main

import (
	"log"

	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_group"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/asset_sub_group"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/auth"
	"github.com/aulaleslie/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, c)
	asset_group.RegisterRoutes(r, c, &authSvc)
	asset_sub_group.RegisterRoutes(r, c, &authSvc)

	r.Run(c.Port)
}
