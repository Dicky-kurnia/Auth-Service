package service

import (
	"service-producer/entity"
	"service-producer/model"
	"service-producer/repository"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type submitSurveyServiceImpl struct {
	repo repository.SubmitSurveyRepository
}

func NewSubmitSurveyService(repo repository.SubmitSurveyRepository) SubmitSurveyService {
	return &submitSurveyServiceImpl{repo: repo}
}

func (s *submitSurveyServiceImpl) SubmitSurvey(request model.SubmitSurveyRequest) error {
	entiUser := entity.SubmitSurvey{
		Id:                                  primitive.NewObjectID(),
		AgentId:                             request.AgentId,
		Age:                                 request.Age,
		Residence:                           request.Residence,
		Gender:                              request.Gender,
		KiosType:                            request.KiosType,
		BusinessAge:                         request.BusinessAge,
		DailyCostumerAverage:                request.DailyCostumerAverage,
		CostumerTypes:                       strings.Split(request.CostumerTypes, ","),
		TopindokuProductsTransaction:        request.TopindokuProductsTransaction,
		TopindokuProductsExpensiveThanOther: strings.Split(request.TopindokuProductsExpensiveThanOther, ","),
		TopindokuProductsOftenTrouble:       strings.Split(request.TopindokuProductsOftenTrouble, ","),
		WhyChooseTopindoku:                  request.WhyChooseTopindoku,
		SatisfactionToTopindokuApps:         request.SatisfactionToTopindokuApps,
		WillRecommendTopindokuApps:          request.WillRecommendTopindokuApps,
		BenefitsFromTopindokuApps:           request.BenefitsFromTopindokuApps,
		HowEasyTopindokuAppsFeatures:        request.HowEasyTopindokuAppsFeatures,
		IsUseOtherSimilarApps:               request.IsUseOtherSimilarApps,
		OtherSimilarAppsNameThatYouUse:      request.OtherSimilarAppsNameThatYouUse,
		HowManyProductsInThatSimilarApps:    request.HowManyProductsInThatSimilarApps,
		ProductsPriceThatCheaperInThatSimilarApps: request.ProductsPriceThatCheaperInThatSimilarApps,
		BenefitsFromThatSimilarApps:               request.BenefitsFromThatSimilarApps,
		IsSatisfiedWithOurCsResponse:              request.IsSatisfiedWithOurCsResponse,
		IsOurCsGaveClearInformation:               request.IsOurCsGaveClearInformation,
		WillAllowOurToContactAboutThisSurvey:      request.WillAllowOurToContactAboutThisSurvey,
		PhoneNumber:                               request.PhoneNumber,
		YourComplaint:                             request.YourComplaint,
	}
	return s.repo.SaveSurvey(entiUser)
}

func (s *submitSurveyServiceImpl) GetReferencesValue() (model.GetRefResponse, error) {
	ref, err := s.repo.GetReferences()
	if err != nil {
		return model.GetRefResponse{}, err
	}
	return model.GetRefResponse{
		Gender:                               ref.Gender,
		KiosType:                             ref.KiosType,
		BusinessAge:                          ref.BusinessAge,
		DailyCustomerAverage:                 ref.DailyCustomerAverage,
		CustomerTypes:                        ref.CustomerTypes,
		TopindokuProductsTransaction:         ref.TopindokuProductsTransaction,
		TopindokuProductsExpensiveThanOther:  ref.TopindokuProductsExpensiveThanOther,
		TopindokuProductsOftenTrouble:        ref.TopindokuProductsOftenTrouble,
		SatisfactionToTopindokuApps:          ref.SatisfactionToTopindokuApps,
		HowEasyTopindokuAppsFeatures:         ref.HowEasyTopindokuAppsFeatures,
		IsUseOtherSimilarApps:                ref.IsUseOtherSimilarApps,
		HowManyProductsInThatSimilarApps:     ref.HowManyProductsInThatSimilarApps,
		IsSatisfiedWithOurCsResponse:         ref.IsSatisfiedWithOurCsResponse,
		IsOurCsGaveClearInformation:          ref.IsOurCsGaveClearInformation,
		WillAllowOurToContactAboutThisSurvey: ref.WillAllowOurToContactAboutThisSurvey,
	}, nil
}
