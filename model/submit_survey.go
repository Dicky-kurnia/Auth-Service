package model

type SubmitSurvey struct {
	AgentId              string `json:"agent_id"`
	Age                  int8   `json:"age"`
	Residence            string `json:"residence"`
	Gender               string `json:"gender"`
	KiosType             string `json:"kios_type"`
	BusinessAge          string `json:"business_age"`
	DailyCostumerAverage string `json:"daily_customer_average"`
	CostumerTypes        string `json:"customer_types"`
	//	CostumerTypes                             []string `form:"costumer_types" json:"costumer_types"`
	TopindokuProductsTransaction        string `json:"topindoku_products_transaction"`
	TopindokuProductsExpensiveThanOther string `json:"topindoku_products_expensive_than_other"`
	//	TopindokuProductsExpensiveThanOther       []string `form:"topindoku_products_expensive_than_other" json:"topindoku_products_expensive_than_other"`
	TopindokuProductsOftenTrouble string `json:"topindoku_products_often_trouble"`
	//	TopindokuProductsOftenTrouble             []string `form:"topindoku_products_often_trouble" json:"topindoku_products_often_trouble"`
	WhyChooseTopindoku                        string `json:"why_choose_topindoku"`
	SatisfactionToTopindokuApps               string `json:"satisfaction_to_topindoku_apps"`
	WillRecommendTopindokuApps                string `json:"will_recommend_topindoku_apps"`
	BenefitsFromTopindokuApps                 string `json:"benefits_from_topindoku_apps"`
	HowEasyTopindokuAppsFeatures              string `json:"how_easy_topindoku_apps_features"`
	IsUseOtherSimilarApps                     string `json:"is_use_other_similar_apps"`
	OtherSimilarAppsNameThatYouUse            string `fjson:"other_similar_apps_name_that_you_use"`
	HowManyProductsInThatSimilarApps          string `json:"how_many_products_in_that_similar_apps"`
	ProductsPriceThatCheaperInThatSimilarApps string `json:"products_price_that_cheaper_in_that_similar_apps"`
	BenefitsFromThatSimilarApps               string `json:"benefits_from_that_similar_apps"`
	IsSatisfiedWithOurCsResponse              string `json:"is_satisfied_with_our_cs_response"`
	IsOurCsGaveClearInformation               string `json:"is_our_cs_gave_clear_information"`
	WillAllowOurToContactAboutThisSurvey      string `json:"will_allow_our_to_contact_about_this_survey"`
	PhoneNumber                               string ` json:"phone_number"`
	YourComplaint                             string `json:"your_complaint"`
}
