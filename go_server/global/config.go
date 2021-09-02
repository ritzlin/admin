package global

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic("failed to map config file: " + err.Error())
	}

	connection, ok := os.LookupEnv(Config.Database.EnvName)
	if !ok {
		panic("failed to find db connection from env: " + Config.Database.EnvName)
	}
	Config.Database.Connection = connection
	fmt.Println("database: " + Config.Database.Connection)
}

type config struct {
	Database database `mapstructure:"database"`
	Log      log      `mapstructure:"log"`
	Server   server   `mapstructure:"server"`
}

type database struct {
	Connection string `mapstructure:"connection"`
	EnvName    string `mapstructure:"env-name"`
}

type log struct {
	Dir   string `mapstructure:"dir"`
	Level string `mapstructure:"level"`
}

type server struct {
	Port int `mapstructure:"port"`
}
