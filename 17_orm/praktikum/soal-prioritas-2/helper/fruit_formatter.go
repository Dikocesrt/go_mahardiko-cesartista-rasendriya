package helper

import "gorm-soal-prioritas-2/model"

func ConvertFruitToJSONFormat(fruit *model.Fruit) interface{} {
    return map[string]interface{}{
        "id":    fruit.ID,
        "name":  fruit.Nama,
        "price": fruit.Harga,
        "nutritions": model.Nutritions{
            Calories:      fruit.Kalori,
            Fat:           fruit.Lemak,
            Sugar:         fruit.Gula,
            Carbohydrates: fruit.Karbohidrat,
            Protein:       fruit.Protein,
        },
    }
}