package config

import "gorm.io/gorm"

type Config struct {
	DB    *gorm.DB `json:"-" gorm:"-"`
	Name  string
	Value string
}

func (Config) TableName() string {
	return "maycms_config"
}
