package repository

import (
	"context"
	"service-producer/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type submitSurveyRepository struct {
	Mongo mongo.Database
}

func NewRepository(Mongo *mongo.Database) SubmitSurveyRepository {
	return &submitSurveyRepository{Mongo: *Mongo}
}

var ctx = context.Background()

func (repo submitSurveyRepository) GetReferences() (entity.SurveyReference, error) {
	surveyReferenceCollection := repo.Mongo.Collection(entity.SurveyReference{}.CollectionName())

	entitySurveyReference := entity.SurveyReference{}
	csr := surveyReferenceCollection.FindOne(ctx, bson.D{})

	err := csr.Decode(&entitySurveyReference)

	return entitySurveyReference, err
}

func (repo submitSurveyRepository) SaveSurvey(survey entity.SubmitSurvey) error {
	userSurveyCollection := repo.Mongo.Collection(entity.SubmitSurvey{}.CollectionName())

	_, err := userSurveyCollection.InsertOne(ctx, survey)
	if err != nil {
		return err
	}
	return nil
}
