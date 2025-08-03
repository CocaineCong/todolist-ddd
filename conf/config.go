package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server *Server `yaml:"service"`
	MySQL  *MySQL  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Port      string `yaml:"port"`
	Version   string `yaml:"version"`
	JwtSecret string `yaml:"jwtSecret"`
	Metrics   string `yaml:"metrics"`
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDbName   int    `yaml:"redisDbName"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
