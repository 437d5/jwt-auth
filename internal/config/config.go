package config

type Config struct {
	Token JWTToken
}

type JWTToken struct {
	Secret string `yaml:"secret"`
}
