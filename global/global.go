package global

import "MottoGo/models"

var (
	// 配置文件
	Configs models.Config
	// 端口
	Port string
	// admin的密码
	KeyAdmin []string
	// user的密码
	KeyUser []string
	// 所有密码用于区分权限
	KeyAll []string
	// 跨域
	AllowCors bool
	// 每秒恢复
	LimitRate uint8
	// 总量
	LimitBurst uint8
	// 基础用户验证
	RequireUserkey bool
)

// // CategoryMap 映射
// var CategoryMap = map[string]string{
// 	"动漫": "a",
// 	"游戏": "b",
// 	"文学": "c",
// 	"原创": "e",
// 	"佚名": "f",
// 	"其他": "g",
// 	"诗句": "h",
// }
