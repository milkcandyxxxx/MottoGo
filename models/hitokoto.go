package models

// Hitokoto 句子json
type Hitokoto struct {
	Id       int    `json:"id"`
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
	From_who string `json:"from_who"`
	Type     string `json:"type"`
}

// Config Config_seve 配置文件json
type Config struct {
	Server   Config_sever `yaml:"server"`
	Security Security     `yaml:"security"`
}
type Config_sever struct {
	Port string `yaml:"port"`
}
type Security struct {
	Key Key `yaml:"key"`
}
type Key struct {
	Admin []string `yaml:"admin"`
	User  []string `yaml:"user"`
}
