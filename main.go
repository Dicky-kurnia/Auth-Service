package main

import (
	"os"
	"service-producer/config"
	"service-producer/controller"
	"service-producer/repository"
	"service-producer/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
func main() {
	mongo := config.MongoConnection()

	//submit Survey
	submitSurveyRepository := repository.NewRepository(mongo)
	submitSurveyService := service.NewSubmitSurveyService(submitSurveyRepository)
	submitSurveyController := controller.NewUserSurveyController(submitSurveyService)

	utilitiController := controller.NewUtilityController(submitSurveyService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1")
	submitSurveyController.Route(v1)
	utilitiController.Route(v1)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
