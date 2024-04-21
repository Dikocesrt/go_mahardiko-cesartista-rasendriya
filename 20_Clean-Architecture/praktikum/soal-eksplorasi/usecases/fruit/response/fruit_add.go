package response

type Fruit struct {
	Name       string    `json:"name" form:"name"`
	ID         int       `json:"id" form:"id"`
	Price      int       `json:"price" form:"price"`
	Nutritions Nutrition `json:"nutritions" form:"nutritions"`
}

type Nutrition struct {
	Calories      float64 `json:"calories" form:"calories"`
	Fat           float64 `json:"fat" form:"fat"`
	Sugar         float64 `json:"sugar" form:"sugar"`
	Carbohydrates float64 `json:"carbohydrates" form:"carbohydrates"`
	Protein       float64 `json:"protein" form:"protein"`
}