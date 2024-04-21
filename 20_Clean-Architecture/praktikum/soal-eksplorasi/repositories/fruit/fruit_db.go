package fruit

import (
	fruitEntities "clean2/entities/fruit"

	"gorm.io/gorm"
)

type Fruit struct {
	gorm.Model
	ID          string     `json:"id" form:"id" gorm:"primary_key"`
	Name        string  `json:"name" form:"name"`
	Price       int     `json:"price" form:"price"`
	Calories      float64 `json:"calories" form:"calories"`
    Fat           float64 `json:"fat" form:"fat"`
    Sugar         float64 `json:"sugar" form:"sugar"`
    Carbohydrates float64 `json:"carbohydrates" form:"carbohydrates"`
    Protein       float64 `json:"protein" form:"protein"`
}

func fromFruitEntitiesToFruitDB(fruit fruitEntities.Fruit) *Fruit {
	return &Fruit{
		ID: fruit.ID,
		Name: fruit.Name,
		Price: fruit.Price,
		Calories: fruit.Nutritions.Calories,
		Fat: fruit.Nutritions.Fat,
		Sugar: fruit.Nutritions.Sugar,
		Carbohydrates: fruit.Nutritions.Carbohydrates,
		Protein: fruit.Nutritions.Protein,
	}
}

func (fruit *Fruit) fromFruitDBToFruitEntities() *fruitEntities.Fruit {
	return &fruitEntities.Fruit{
		ID: fruit.ID,
		Name: fruit.Name,
		Price: fruit.Price,
		Nutritions: fruitEntities.Nutrition{
			Calories: fruit.Calories,
			Fat: fruit.Fat,
			Sugar: fruit.Sugar,
			Carbohydrates: fruit.Carbohydrates,
			Protein: fruit.Protein,
		},
	}
}