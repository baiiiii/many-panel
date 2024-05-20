package request

import "github.com/1Panel-dev/1Panel/backend/app/dto"

type MpHostSearch struct {
	dto.PageInfo
	Name string `json:"name"`
}

type CreateMpHost struct {
	Name string `json:"name" form:"name" gorm:"comment:主机名" validate:"required"`  //主机名
	Host string `json:"host" form:"host" gorm:"comment:主机地址" validate:"required"` //主机地址
	User string `json:"user" form:"user" gorm:"comment:用户名" validate:"required"`  //用户名
	Pwd  string `json:"pwd" form:"pwd" gorm:"comment:密码" validate:"required"`     //密码
}

type UpdateMpHost struct {
	ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" gorm:"comment:主机名" validate:"required"`  //主机名
	Host string `json:"host" form:"host" gorm:"comment:主机地址" validate:"required"` //主机地址
	User string `json:"user" form:"user" gorm:"comment:用户名" validate:"required"`  //用户名
	Pwd  string `json:"pwd" form:"pwd" gorm:"comment:密码" validate:"required"`     //密码
}

type DeleteMpHost struct {
	Ids []uint `json:"ids" validate:"required"`
}
