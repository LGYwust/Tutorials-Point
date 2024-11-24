package config

import "fmt"

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Config       string `yaml:"config"`
	DB           string `yaml:"db"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `yaml:"log-mode"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DB, m.Config)
}
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
