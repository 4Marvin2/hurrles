package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host     string
	HttpPort string
	GrpcUrl  string
	CertFile string
	KeyFile  string
}

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
	Hurrles1      ServerConfig
	Hurrles2      ServerConfig
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

	Hurrles1 = ServerConfig{
		HttpPort: viper.GetString(`hurrles_1.httpPort`),
		Host:     viper.GetString(`hurrles_1.host`),
		CertFile: viper.GetString(`hurrles_1.certFile`),
		KeyFile:  viper.GetString(`hurrles_1.keyFile`),
	}

	Hurrles2 = ServerConfig{
		HttpPort: viper.GetString(`hurrles_2.httpPort`),
		Host:     viper.GetString(`hurrles_2.host`),
		CertFile: viper.GetString(`hurrles_2.certFile`),
		KeyFile:  viper.GetString(`hurrles_2.keyFile`),
	}

	Tarantool = TarantoolConfig{
		Port:     viper.GetString(`tarantool.port`),
		Host:     viper.GetString(`tarantool.host`),
		User:     viper.GetString(`tarantool.user`),
		Password: viper.GetString(`tarantool.pass`),
		DBName:   viper.GetString(`tarantool.name`),
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
