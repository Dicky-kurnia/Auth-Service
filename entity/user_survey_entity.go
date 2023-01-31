package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubmitSurvey struct {
	Id                                        primitive.ObjectID `bson:"_id" json:"_id"`
	AgentId                                   string             `bson:"agent_id" json:"agent_id"`
	Age                                       int8               `bson:"age" json:"age"`
	Residence                                 string             `bson:"residence" json:"residence"`
	Gender                                    string             `bson:"gender" json:"gender"`
	KiosType                                  string             `bson:"kios_type" json:"kios_type"`
	BusinessAge                               string             `bson:"business_age" json:"business_age"`
	DailyCostumerAverage                      string             `bson:"daily_costumer_average" json:"daily_costumer_average"`
	CostumerTypes                             []string           `bson:"costumer_types" json:"costumer_types"`
	TopindokuProductsTransaction              string             `bson:"topindoku_products_transaction" json:"topindoku_products_transaction"`
	TopindokuProductsExpensiveThanOther       []string           `bson:"topindoku_products_expensive_than_other" json:"topindoku_products_expensive_than_other"`
	TopindokuProductsOftenTrouble             []string           `bson:"topindoku_products_often_trouble" json:"topindoku_products_often_trouble"`
	WhyChooseTopindoku                        string             `bson:"why_choose_topindoku" json:"why_choose_topindoku"`
	SatisfactionToTopindokuApps               string             `bson:"satisfaction_to_topindoku_apps" json:"satisfaction_to_topindoku_apps"`
	WillRecommendTopindokuApps                string             `bson:"will_recommend_topindoku_apps" json:"will_recommend_topindoku_apps"`
	BenefitsFromTopindokuApps                 string             `bson:"benefits_from_topindoku_apps" json:"benefits_from_topindoku_apps"`
	HowEasyTopindokuAppsFeatures              string             `bson:"how_easy_topindoku_apps_features" json:"how_easy_topindoku_apps_features"`
	IsUseOtherSimilarApps                     string             `bson:"is_use_other_similar_apps" json:"is_use_other_similar_apps"`
	OtherSimilarAppsNameThatYouUse            string             `bson:"other_similar_apps_name_that_you_use" json:"other_similar_apps_name_that_you_use"`
	HowManyProductsInThatSimilarApps          string             `bson:"how_many_products_in_that_similar_apps" json:"how_many_products_in_that_similar_apps"`
	ProductsPriceThatCheaperInThatSimilarApps string             `bson:"products_price_that_cheaper_in_that_similar_apps" json:"products_price_that_cheaper_in_that_similar_apps"`
	BenefitsFromThatSimilarApps               string             `bson:"benefits_from_that_similar_apps" json:"benefits_from_that_similar_apps"`
	IsSatisfiedWithOurCsResponse              string             `bson:"is_satisfied_with_our_cs_response" json:"is_satisfied_with_our_cs_response"`
	IsOurCsGaveClearInformation               string             `bson:"is_our_cs_gave_clear_information" json:"is_our_cs_gave_clear_information"`
	WillAllowOurToContactAboutThisSurvey      string             `bson:"will_allow_our_to_contact_about_this_survey" json:"will_allow_our_to_contact_about_this_survey"`
	PhoneNumber                               string             `bson:"phone_number" json:"phone_number"`
	YourComplaint                             string             `bson:"your_complaint" json:"your_complaint"`
}

func (SubmitSurvey) CollectionName() string {
	return "user_survey"
}
