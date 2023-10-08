package config

import "os"

type Config struct {
	RDB struct {
		Host     string
		Port     string
		UserName string
		Password string
		Database string
	}
}

func NewConfig() *Config {
	conf := new(Config)

	conf.RDB.Host = os.Getenv("DB_HOST")
	conf.RDB.Port = os.Getenv("DB_PORT")
	conf.RDB.UserName = os.Getenv("DB_USER")
	conf.RDB.Password = os.Getenv("DB_PASSWORD")
	conf.RDB.Database = os.Getenv("DB_DATABASE")

	return conf
}
