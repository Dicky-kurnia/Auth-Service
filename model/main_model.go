package model

import "service-producer/entity"

type SubmitSurveyRequest struct {
	AgentId              string `json:"agent_id"`
	Age                  int8   `form:"age" json:"age"`
	Residence            string `form:"residence" json:"residence"`
	Gender               string `form:"gender" json:"gender"`
	KiosType             string `form:"kios_type" json:"kios_type"`
	BusinessAge          string `form:"business_age" json:"business_age"`
	DailyCostumerAverage string `form:"daily_customer_average" json:"daily_customer_average"`
	CostumerTypes        string `form:"customer_types" json:"customer_types"`
	//	CostumerTypes                             []string `form:"costumer_types" json:"costumer_types"`
	TopindokuProductsTransaction        string `form:"topindoku_products_transaction" json:"topindoku_products_transaction"`
	TopindokuProductsExpensiveThanOther string `form:"topindoku_products_expensive_than_other" json:"topindoku_products_expensive_than_other"`
	//	TopindokuProductsExpensiveThanOther       []string `form:"topindoku_products_expensive_than_other" json:"topindoku_products_expensive_than_other"`
	TopindokuProductsOftenTrouble string `form:"topindoku_products_often_trouble" json:"topindoku_products_often_trouble"`
	//	TopindokuProductsOftenTrouble             []string `form:"topindoku_products_often_trouble" json:"topindoku_products_often_trouble"`
	WhyChooseTopindoku                        string `form:"why_choose_topindoku" json:"why_choose_topindoku"`
	SatisfactionToTopindokuApps               string `form:"satisfaction_to_topindoku_apps" json:"satisfaction_to_topindoku_apps"`
	WillRecommendTopindokuApps                string `form:"will_recommend_topindoku_apps" json:"will_recommend_topindoku_apps"`
	BenefitsFromTopindokuApps                 string `form:"benefits_from_topindoku_apps" json:"benefits_from_topindoku_apps"`
	HowEasyTopindokuAppsFeatures              string `form:"how_easy_topindoku_apps_features" json:"how_easy_topindoku_apps_features"`
	IsUseOtherSimilarApps                     string `form:"is_use_other_similar_apps" json:"is_use_other_similar_apps"`
	OtherSimilarAppsNameThatYouUse            string `form:"other_similar_apps_name_that_you_use" json:"other_similar_apps_name_that_you_use"`
	HowManyProductsInThatSimilarApps          string `form:"how_many_products_in_that_similar_apps" json:"how_many_products_in_that_similar_apps"`
	ProductsPriceThatCheaperInThatSimilarApps string `form:"products_price_that_cheaper_in_that_similar_apps" json:"products_price_that_cheaper_in_that_similar_apps"`
	BenefitsFromThatSimilarApps               string `form:"benefits_from_that_similar_apps" json:"benefits_from_that_similar_apps"`
	IsSatisfiedWithOurCsResponse              string `form:"is_satisfied_with_our_cs_response" json:"is_satisfied_with_our_cs_response"`
	IsOurCsGaveClearInformation               string `form:"is_our_cs_gave_clear_information" json:"is_our_cs_gave_clear_information"`
	WillAllowOurToContactAboutThisSurvey      string `form:"will_allow_our_to_contact_about_this_survey" json:"will_allow_our_to_contact_about_this_survey"`
	PhoneNumber                               string `form:"phone_number" json:"phone_number"`
	YourComplaint                             string `form:"your_complaint" json:"your_complaint"`
}

type GetRefResponse struct {
	Gender                               []entity.LabelValue `json:"gender"`
	KiosType                             []entity.LabelValue `json:"kios_type"`
	BusinessAge                          []entity.LabelValue `json:"business_age"`
	DailyCustomerAverage                 []entity.LabelValue `json:"daily_customer_average"`
	CustomerTypes                        []entity.LabelValue `json:"customer_types"`
	TopindokuProductsTransaction         []entity.LabelValue `json:"topindoku_products_transaction"`
	TopindokuProductsExpensiveThanOther  []entity.LabelValue `json:"topindoku_products_expensive_than_other"`
	TopindokuProductsOftenTrouble        []entity.LabelValue `json:"topindoku_products_often_trouble"`
	SatisfactionToTopindokuApps          []entity.LabelValue `json:"satisfaction_to_topindoku_apps"`
	HowEasyTopindokuAppsFeatures         []entity.LabelValue `json:"how_easy_topindoku_apps_features"`
	IsUseOtherSimilarApps                []entity.LabelValue `json:"is_use_other_similar_apps"`
	HowManyProductsInThatSimilarApps     []entity.LabelValue `json:"how_many_products_in_that_similar_apps"`
	IsSatisfiedWithOurCsResponse         []entity.LabelValue `json:"is_satisfied_with_our_cs_response"`
	IsOurCsGaveClearInformation          []entity.LabelValue `json:"is_our_cs_gave_clear_information"`
	WillAllowOurToContactAboutThisSurvey []entity.LabelValue `json:"will_allow_our_to_contact_about_this_survey"`
}

type RegistrerUserRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required, email"`
}
