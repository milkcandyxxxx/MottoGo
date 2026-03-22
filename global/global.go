package global

import "MottoGo/models"

var (
	// Hit      []models.Hitokoto
	Configs  models.Config
	Port     string
	KeyAdmin []string
	KeyUser  []string
	KeyAll   []string
	Hit      map[string][]models.Hitokoto
	AllHit   []models.Hitokoto
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
