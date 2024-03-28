package model

type Fruit struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Nama        string  `json:"name"`
	Harga       int     `json:"price"`
	Kalori      float64 `json:"calories"`
	Lemak       float64 `json:"fat"`
	Gula        float64 `json:"sugar"`
	Karbohidrat float64 `json:"carbohydrates"`
	Protein     float64 `json:"protein"`
}

type Nutritions struct {
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}
