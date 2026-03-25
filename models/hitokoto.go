package models

import "time"

type Hitokoto struct {
	// ID: 自增主键，JSON 序列化时隐藏
	ID uint `gorm:"primaryKey;autoIncrement" json:"-"`
	// Uuid: 唯一标识，固定长度索引
	Uuid string `gorm:"uniqueIndex;not null;size:36" json:"uuid"`
	// Hitokoto: 正文内容，长文本
	Hitokoto string `gorm:"type:text;not null" json:"hitokoto"`
	// Source: 来源，默认“未知”索引
	Source string `gorm:"default:'未知';index" json:"source"`
	// Author: 作者，默认“佚名”索引
	Author string `gorm:"default:'佚名';index" json:"author"`
	// Category: 分类标签，非空索引
	Category string `gorm:"index;not null;size:20" json:"category"`
	// CreatedAt: 自动创建时间戳
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Config 配置文件
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
