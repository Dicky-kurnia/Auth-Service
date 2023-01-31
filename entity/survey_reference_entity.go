package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type LabelValue struct {
	Label string `bson:"label"  json:"label"`
	Value string `bson:"value" json:"value"`
}

type SurveyReference struct {
	Id                                   primitive.ObjectID `bson:"_id" json:"_id"`
	Gender                               []LabelValue       `bson:"gender" json:"gender"`
	KiosType                             []LabelValue       `bson:"kios_type" json:"kios_type"`
	BusinessAge                          []LabelValue       `bson:"business_age" json:"business_age"`
	DailyCustomerAverage                 []LabelValue       `bson:"daily_customer_average" json:"daily_customer_average"`
	CustomerTypes                        []LabelValue       `bson:"customer_types" json:"customer_types"`
	TopindokuProductsTransaction         []LabelValue       `bson:"topindoku_products_transaction" json:"topindoku_products_transaction"`
	TopindokuProductsExpensiveThanOther  []LabelValue       `bson:"topindoku_products_expensive_than_other" json:"topindoku_products_expensive_than_other"`
	TopindokuProductsOftenTrouble        []LabelValue       `bson:"topindoku_products_often_trouble" json:"topindoku_products_often_trouble"`
	SatisfactionToTopindokuApps          []LabelValue       `bson:"satisfaction_to_topindoku_apps" json:"satisfaction_to_topindoku_apps"`
	HowEasyTopindokuAppsFeatures         []LabelValue       `bson:"how_easy_topindoku_apps_features" json:"how_easy_topindoku_apps_features"`
	IsUseOtherSimilarApps                []LabelValue       `bson:"is_use_other_similar_apps" json:"is_use_other_similar_apps"`
	HowManyProductsInThatSimilarApps     []LabelValue       `bson:"how_many_products_in_that_similar_apps" json:"how_many_products_in_that_similar_apps"`
	IsSatisfiedWithOurCsResponse         []LabelValue       `bson:"is_satisfied_with_our_cs_response" json:"is_satisfied_with_our_cs_response"`
	IsOurCsGaveClearInformation          []LabelValue       `bson:"is_our_cs_gave_clear_information" json:"is_our_cs_gave_clear_information"`
	WillAllowOurToContactAboutThisSurvey []LabelValue       `bson:"will_allow_our_to_contact_about_this_survey" json:"will_allow_our_to_contact_about_this_survey"`
}

func (SurveyReference) CollectionName() string {
	return "survey_reference"
}
