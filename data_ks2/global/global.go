package global

import (
	"data_ks2/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
