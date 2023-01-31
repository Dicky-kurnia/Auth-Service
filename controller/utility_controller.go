package controller

import (
	"fmt"
	"os"
	"service-producer/exception"
	"service-producer/helper"
	"service-producer/model"
	"service-producer/service"

	"github.com/gofiber/fiber/v2"
)

type utilityController struct {
	service service.SubmitSurveyService
}

func NewUtilityController(service service.SubmitSurveyService) *utilityController {
	return &utilityController{service: service}
}
func (controller utilityController) Route(app fiber.Router) {

	router := app.Group("/utility")
	router.Get("/ref-value", controller.GetRefValue)
	router.Get("/error-code-list", controller.GetErrorCode)

	if os.Getenv("IS_DEVEL") == "1" {
		router.Get("/create-new-token", controller.CreateNewToken)
	}
}
func (controller utilityController) GetRefValue(c *fiber.Ctx) error {
	response, err := controller.service.GetReferencesValue()
	exception.PanicIfNeeded(err)
	return c.JSON(model.Response{
		Code:   200,
		Status: model.OK,
		Data:   response,
	})
}
func (controller utilityController) GetErrorCode(c *fiber.Ctx) error {
	return c.JSON(model.Response{
		Code:   200,
		Status: model.OK,
		Data: map[string]string{
			"AGENT_BLOCKED":   "Akun ini telah terblokir. Silahkan hubungi customer service",
			"AGENT_INACTIVE":  "Akun ini tidak aktif. Silahkan hubungi customer service",
			"AGENT_NOT_FOUND": "Data user tidak ditemukan",
		},
	})
}
func (controller utilityController) CreateNewToken(c *fiber.Ctx) error {
	agentId := c.Query("agent_id")
	if agentId == "" {
		panic(exception.ValidationError{
			Message: fmt.Sprintf(`{"%s": "%s"}`, "agent_id", model.NOT_BLANK_ERR_TYPE),
		})
	}
	td := helper.CreateToken(model.JwtPayload{
		AgentId: agentId,
	})
	return c.JSON(model.Response{
		Code:   200,
		Status: model.OK,
		Data: map[string]string{
			"access_token": td.AccessToken,
		},
	})
}
