package models

import (
	"time"
)

// UniversalProducts model
type UniversalProducts struct {
	Results []struct {
		ID                        string             `json:"id"`
		DisplayName               string             `json:"display_name"`
		ProductName               string             `json:"product_name"`
		BrandName                 string             `json:"brand_name"`
		Description               string             `json:"description"`
		CommonUpc                 string             `json:"common_upc"`
		UnitOfMeasure             string             `json:"unit_of_measure"`
		DeletedAt                 time.Time          `json:"deleted_at"`
		QaComplete                bool               `json:"qa_complete"`
		Size                      string             `json:"size"`
		ProductType               string             `json:"product_type"`
		UpdatedAt                 time.Time          `json:"updated_at"`
		CreatedAt                 time.Time          `json:"created_at"`
		Keywords                  string             `json:"keywords"`
		AllergenInfo              string             `json:"allergen_info"`
		Prop65                    string             `json:"prop_65"`
		ChokingHazard             string             `json:"choking_hazard"`
		Ingredients               string             `json:"ingredients"`
		IsAlcohol                 string             `json:"is_alcohol"`
		ProductWarnings           string             `json:"product_warnings"`
		ProductPoundage           float64            `json:"product_poundage"`
		ProductPoundageConfidence float64            `json:"product_poundage_confidence"`
		ProductIds                []int              `json:"product_ids"`
		NutritionalDetails        NutritionalDetails `json:"nutritional_details"`
	}
}

// NutritionalDetails does something
type NutritionalDetails struct {
	Products []struct {
		Name   string `json:"name"`
		Labels []struct {
			Title   string
			Serving struct {
				Percontainer string `json:"perContainer"`
				Size         string `json:"size"`
			} `json:"serving"`
			Calories struct {
				Total string `json:"total"`
				Fat   string `json:"fat"`
			} `json:"calories"`
			Fat struct {
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
				Trans struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"trans"`
				Saturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"saturated"`
				Polyunsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"polyunsaturated"`
				Monounsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"monounsaturated"`
			}
			Carbs struct {
				Sugar struct {
					Total struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"total"`
					Added struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"added"`
					Alcohol struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"alcohol"`
				} `json:"sugar"`
				Fiber struct {
					Dietary struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"dietary"`
					Soluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"soluable"`
					Insoluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"Insoluable"`
				} `json:"fiber"`
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
			} `json:"carbs"`
			Cholesterol struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"cholesterol"`
			Sodium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"sodium"`
			Protein struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"protein"`
			Vitamina struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminA"`
			Vitaminc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminC"`
			Calcium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"calcium"`
			Iron struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iron"`
			Vitamind struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminD"`
			Vitamine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminE"`
			Vitamink struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminK"`
			Thiamin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"thiamin"`
			Riboflavin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"riboflavin"`
			Niacin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"niacin"`
			Vitaminb6 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB6"`
			Folate struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folate"`
			Folicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folicAcid"`
			Vitaminb12 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB12"`
			Biotin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"biotin"`
			Pantothenicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"pantothenicAcid"`
			Phosphorus struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"phosphorus"`
			Iodine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iodine"`
			Magnesium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"magnesium"`
			Zinc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"zinc"`
			Selenium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"selenium"`
			Copper struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"copper"`
			Manganese struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"manganese"`
			Chromium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chromium"`
			Molybdenum struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"molybdenum"`
			Chloride struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chloride"`
			Potassium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"potassium"`
			Choline struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"choline"`
			Proteinnutrients struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"proteinNutrients"`
			Nutrients []struct {
				Name     string `json:"name"`
				Uom      string `json:"uom"`
				Value    string `json:"value"`
				Dailypct string `json:"dailyPct"`
			} `json:"nutrients"`
		} `json:"labels"`
		Ingredients string `json:"ingredients"`
		Warnings    string `json:"warnings"`
		Audience    string `json:"audience"`
	} `json:"products"`
}

// MalformedProducts does something
type MalformedProducts struct {
	Products struct {
		Name   string `json:"name"`
		Labels []struct {
			Title   string
			Serving struct {
				Percontainer string `json:"perContainer"`
				Size         string `json:"size"`
			} `json:"serving"`
			Calories struct {
				Total string `json:"total"`
				Fat   string `json:"fat"`
			} `json:"calories"`
			Fat struct {
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
				Trans struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"trans"`
				Saturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"saturated"`
				Polyunsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"polyunsaturated"`
				Monounsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"monounsaturated"`
			} `json:"fat"`
			Carbs struct {
				Sugar struct {
					Total struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"total"`
					Added struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"added"`
					Alcohol struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"alcohol"`
				} `json:"sugar"`
				Fiber struct {
					Dietary struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"dietary"`
					Soluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"soluable"`
					Insoluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"Insoluable"`
				} `json:"fiber"`
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
			} `json:"carbs"`
			Cholesterol struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"cholesterol"`
			Sodium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"sodium"`
			Protein struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"protein"`
			Vitamina struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminA"`
			Vitaminc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminC"`
			Calcium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"calcium"`
			Iron struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iron"`
			Vitamind struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminD"`
			Vitamine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminE"`
			Vitamink struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminK"`
			Thiamin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"thiamin"`
			Riboflavin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"riboflavin"`
			Niacin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"niacin"`
			Vitaminb6 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB6"`
			Folate struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folate"`
			Folicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folicAcid"`
			Vitaminb12 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB12"`
			Biotin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"biotin"`
			Pantothenicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"pantothenicAcid"`
			Phosphorus struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"phosphorus"`
			Iodine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iodine"`
			Magnesium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"magnesium"`
			Zinc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"zinc"`
			Selenium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"selenium"`
			Copper struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"copper"`
			Manganese struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"manganese"`
			Chromium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chromium"`
			Molybdenum struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"molybdenum"`
			Chloride struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chloride"`
			Potassium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"potassium"`
			Choline struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"choline"`
			Proteinnutrients struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"proteinNutrients"`
			Nutrients []struct {
				Name     string `json:"name"`
				Uom      string `json:"uom"`
				Value    string `json:"value"`
				Dailypct string `json:"dailyPct"`
			} `json:"nutrients"`
		} `json:"labels"`
		Ingredients string `json:"ingredients"`
		Warnings    string `json:"warnings"`
		Audience    string `json:"audience"`
	} `json:"products"`
}

// MalformedLabels does something
type MalformedLabels struct {
	Products []struct {
		Name   string `json:"name"`
		Labels struct {
			Title   string
			Serving struct {
				Percontainer string `json:"perContainer"`
				Size         string `json:"size"`
			} `json:"serving"`
			Calories struct {
				Total string `json:"total"`
				Fat   string `json:"fat"`
			} `json:"calories"`
			Fat struct {
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
				Trans struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"trans"`
				Saturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"saturated"`
				Polyunsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"polyunsaturated"`
				Monounsaturated struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"monounsaturated"`
			}
			Carbs struct {
				Sugar struct {
					Total struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"total"`
					Added struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"added"`
					Alcohol struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"alcohol"`
				} `json:"sugar"`
				Fiber struct {
					Dietary struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"dietary"`
					Soluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"soluable"`
					Insoluable struct {
						Value    string `json:"value"`
						Dailypct string `json:"dailyPct"`
					} `json:"Insoluable"`
				} `json:"fiber"`
				Total struct {
					Dailypct string `json:"dailyPct"`
					Value    string `json:"value"`
				} `json:"total"`
			} `json:"carbs"`
			Cholesterol struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"cholesterol"`
			Sodium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"sodium"`
			Protein struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"protein"`
			Vitamina struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminA"`
			Vitaminc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminC"`
			Calcium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"calcium"`
			Iron struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iron"`
			Vitamind struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminD"`
			Vitamine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminE"`
			Vitamink struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminK"`
			Thiamin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"thiamin"`
			Riboflavin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"riboflavin"`
			Niacin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"niacin"`
			Vitaminb6 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB6"`
			Folate struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folate"`
			Folicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"folicAcid"`
			Vitaminb12 struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"vitaminB12"`
			Biotin struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"biotin"`
			Pantothenicacid struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"pantothenicAcid"`
			Phosphorus struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"phosphorus"`
			Iodine struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"iodine"`
			Magnesium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"magnesium"`
			Zinc struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"zinc"`
			Selenium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"selenium"`
			Copper struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"copper"`
			Manganese struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"manganese"`
			Chromium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chromium"`
			Molybdenum struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"molybdenum"`
			Chloride struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"chloride"`
			Potassium struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"potassium"`
			Choline struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"choline"`
			Proteinnutrients struct {
				Dailypct string `json:"dailyPct"`
				Value    string `json:"value"`
			} `json:"proteinNutrients"`
			Nutrients []struct {
				Name     string `json:"name"`
				Uom      string `json:"uom"`
				Value    string `json:"value"`
				Dailypct string `json:"dailyPct"`
			} `json:"nutrients"`
		} `json:"labels"`
		Ingredients string `json:"ingredients"`
		Warnings    string `json:"warnings"`
		Audience    string `json:"audience"`
	} `json:"products"`
}

// QueryMalformedProducts does something
type QueryMalformedProducts struct {
	UniversalID        string
	Nutritionaldetails MalformedProducts
}

// QueryMalformedLabels does something
type QueryMalformedLabels struct {
	UniversalID        string
	Nutritionaldetails MalformedLabels
}
