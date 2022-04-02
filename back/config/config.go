package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type TarantoolConfig struct {
	User     string
	Password string
	Port     string
	Host     string
	DBName   string
}

type PostgresConfig struct {
	User     string
	Password string
	Port     string
	Host     string
	DBName   string
}

type TimeoutsConfig struct {
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	ContextTimeout time.Duration
}

type contextUserID string

type contextUser string

var (
	Tarantool     TarantoolConfig
	Postgres      PostgresConfig
	Timeouts      TimeoutsConfig
	ContextUserID contextUserID
	ContextUser   contextUser
)

func SetConfig() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	Tarantool = TarantoolConfig{
		Port:     viper.GetString(`session.port`),
		Host:     viper.GetString(`session.host`),
		User:     viper.GetString(`session.user`),
		Password: viper.GetString(`session.pass`),
		DBName:   viper.GetString(`session.name`),
	}

	Postgres = PostgresConfig{
		Port:     viper.GetString(`postgres.port`),
		Host:     viper.GetString(`postgres.host`),
		User:     viper.GetString(`postgres.user`),
		Password: viper.GetString(`postgres.pass`),
		DBName:   viper.GetString(`postgres.name`),
	}

	Timeouts = TimeoutsConfig{
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
		ContextTimeout: time.Second * 2,
	}

	ContextUserID = "userID"

	ContextUser = "user"
}
