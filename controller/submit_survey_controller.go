package controller

import (
	"service-producer/exception"
	"service-producer/middleware"
	"service-producer/model"
	"service-producer/service"

	"github.com/gofiber/fiber/v2"
)

type userSurveyController struct {
	service service.SubmitSurveyService
}

func NewUserSurveyController(service service.SubmitSurveyService) *userSurveyController {
	return &userSurveyController{service: service}
}
func (controller userSurveyController) Route(app fiber.Router) {
	router := app.Group("/main")
	router.Post("/submit-survey", middleware.CheckToken(), controller.SubmitUserSurvey)
}
func (controller userSurveyController) SubmitUserSurvey(c *fiber.Ctx) error {
	agentId := c.Locals("currentAgentId").(string)
	request := new(model.SubmitSurveyRequest)

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	request.AgentId = agentId

	err = controller.service.SubmitSurvey(*request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.Response{
		Code:   200,
		Status: model.OK,
	})
}
