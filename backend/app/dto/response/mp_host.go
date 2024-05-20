package response

type MpHost struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"` //主机名
	Host      string `json:"host"` //主机地址
	User      string `json:"user"` //用户名
	IsDefault string `json:"isDefault"`
}
