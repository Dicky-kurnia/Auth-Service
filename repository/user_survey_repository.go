package repository

import "service-producer/entity"

type SubmitSurveyRepository interface {
	SaveSurvey(survey entity.SubmitSurvey) error
	GetReferences() (entity.SurveyReference, error)
}
