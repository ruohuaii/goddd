package server

import (
	"github.com/ruohuaii/goddd/application"
	"github.com/ruohuaii/goddd/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//merchant handler
	merchantRepo := persistence.NewMerchantRepository()
	jwtSvc := application.NewJwtService()
	merchantSvc := application.NewMerchantService(merchantRepo, jwtSvc)
	merchantHandler := NewMerchantHandler(merchantSvc)

	openapi := router.Group("/openapi/v1")
	{
		openapi.POST("/signup", merchantHandler.Signup)
		openapi.POST("/signin", merchantHandler.Signin)
	}

	return router
}
