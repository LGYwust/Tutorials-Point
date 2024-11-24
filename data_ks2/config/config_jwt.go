package config

type Jwt struct {
	Exp    int    `yaml:"exp"`
	Iss    string `yaml:"iss"`
	Secret string `yaml:"secret"`
}
