package config

import "fmt"

type Redis struct {
	IP       string `yaml:"ip"`
	Port     uint   `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolsize"`
}

func (redis Redis) Addr() string {
	return fmt.Sprintf("%s:%d", redis.IP, redis.Port)
}
