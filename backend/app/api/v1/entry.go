package v1

import "github.com/1Panel-dev/1Panel/backend/app/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var (
	authService = service.NewIAuthService()

	settingService = service.NewISettingService()

	logService = service.NewILogService()

	mpHostService = service.NewIMpHostService()
)
