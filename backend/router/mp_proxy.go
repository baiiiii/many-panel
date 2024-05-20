package router

import (
	v1 "github.com/1Panel-dev/1Panel/backend/app/api/v1"
	"github.com/1Panel-dev/1Panel/backend/middleware"
	"github.com/gin-gonic/gin"
)

type MpProxyRouter struct{}

func (s *MpProxyRouter) InitRouter(Router *gin.RouterGroup) {
	mpProxyRouter := Router.Use(middleware.JwtAuth()).
		Use(middleware.SessionAuth()).
		Use(middleware.PasswordExpired())
	baseApi := v1.ApiGroupApp.BaseApi
	mpProxyRouter.Any("/*name", baseApi.ApiProxyHandler)
}
