package main

import (
	"aqary/delivery"
	"aqary/repository"
	"aqary/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	err error
)

func main() {

	config, err := LoadConfig()
	if err != nil {
		log.Println("error reading env:", err)
	}
	server, err := InitApi(config)
	if err != nil {
		log.Println("error init api:", err)
	}
	server.Run(":8080")
}

func InitApi(config Config) (*gin.Engine, error) {
	Db, err := repository.InitDatabase()
	if err != nil {
		return nil, err
	}

	repoLayer := repository.InitRepository(Db)
	usecaseLayer := usecase.InitUsecase(repoLayer)
	handlers := delivery.InitHandlers(usecaseLayer)

	router := gin.Default()

	router = delivery.InitRoutes(router, handlers, &gin.Context{})

	return router, nil

}

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	JwtKey     string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Println("err reading .env viper", err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Println("err unmarshaling config viper", err)
		return
	}

	return
}
