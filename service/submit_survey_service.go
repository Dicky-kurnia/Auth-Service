package service

import "service-producer/model"

type SubmitSurveyService interface {
	SubmitSurvey(request model.SubmitSurveyRequest) error
	GetReferencesValue() (model.GetRefResponse, error)
}
