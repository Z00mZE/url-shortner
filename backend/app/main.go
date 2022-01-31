package main

import (
	"github.com/Z00mZE/url-shortner/config"
	"github.com/Z00mZE/url-shortner/internal/app"
	"log"
)

func main() {
	//	получаем настройки
	configuration, configError := config.NewConfig()
	//	что-ьл пошло не так, смысла продолжать нет
	if configError != nil {
		log.Fatalf("Config error: %s", configError)
	}
	//	запуск сервера приложения
	app.Run(configuration)
}
