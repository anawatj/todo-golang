package config

type Config struct {
	DB *DBConfig
}
type JwtCode struct {
	Key string
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetJwtKey() *JwtCode {
	return &JwtCode{
		Key: "secret",
	}
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "todo",
			Charset:  "utf8",
		},
	}
}
