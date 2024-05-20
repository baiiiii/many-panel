package model

type MpHost struct {
	BaseModel
	Name string `json:"name" form:"name" gorm:"comment:主机名"`
	Host string `json:"host" form:"host" gorm:"comment:主机地址"`
	User string `json:"user" form:"user" gorm:"comment:用户名"`
	Pwd  string `json:"-" form:"pwd" gorm:"comment:密码"`
}
