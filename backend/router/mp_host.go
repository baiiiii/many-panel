package router

import (
	v1 "github.com/1Panel-dev/1Panel/backend/app/api/v1"
	"github.com/1Panel-dev/1Panel/backend/middleware"

	"github.com/gin-gonic/gin"
)

type MpHostRouter struct{}

func (s *MpHostRouter) InitRouter(Router *gin.RouterGroup) {
	mpHostRouter := Router.Group("host").
		Use(middleware.JwtAuth()).
		Use(middleware.SessionAuth()).
		Use(middleware.PasswordExpired())
	baseApi := v1.ApiGroupApp.BaseApi
	{
		mpHostRouter.POST("create", baseApi.CreateMpHost)      // 新建mpHost表
		mpHostRouter.POST("delete", baseApi.DeleteMpHostByIds) // 批量删除mpHost表
		mpHostRouter.POST("update", baseApi.UpdateMpHost)      // 更新mpHost表
		mpHostRouter.POST("search", baseApi.GetMpHostList)
		mpHostRouter.GET("default/:id", baseApi.SetDefaultHost)
		mpHostRouter.GET("login/:id", baseApi.LoginMpHost)
		mpHostRouter.GET("refreshToken/:id", baseApi.LoginMpHost)
	}
}
