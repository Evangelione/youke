package config

import (
	"errors"
	"yk/internal/pkg/infra"

	"gorm.io/gorm"
)

func FindOne(key string) (*Config, error) {
	c := Config{
		DB:   infra.Mysql,
		Name: key,
	}

	if err := c.DB.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("未找到记录")
		}
		return nil, err
	}

	return &c, nil
}
