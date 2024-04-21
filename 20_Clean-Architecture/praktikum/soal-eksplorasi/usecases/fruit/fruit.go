package fruit

import (
	"clean2/constants"
	fruitEntities "clean2/entities/fruit"
	"clean2/usecases/fruit/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type FruitUseCase struct {
	repository fruitEntities.RepositoryInterface
}

func NewFruitUseCase(repository fruitEntities.RepositoryInterface) *FruitUseCase {
	return &FruitUseCase{
		repository: repository,
	}
}

func (fruitUseCase *FruitUseCase) GetAllFruits() ([]fruitEntities.Fruit, error) {
	var fruitsEntities []fruitEntities.Fruit

	fruitsEntities, err := fruitUseCase.repository.GetAllFruits()
	if err != nil {
		return []fruitEntities.Fruit{}, constants.ErrGetDatabase
	}

	return fruitsEntities, nil
}

func (fruitUseCase *FruitUseCase) GetFruitById(id string) (fruitEntities.Fruit, error) {
	if id == "" {
		return fruitEntities.Fruit{}, constants.ErrEmptyInput
	}

	fruitEnt, err := fruitUseCase.repository.GetFruitById(id)
	if err != nil {
		return fruitEntities.Fruit{}, constants.ErrGetDatabase
	}

	return fruitEnt, nil
}

func (fruitUseCase *FruitUseCase) AddFruit(fruitEnt fruitEntities.Fruit) (fruitEntities.Fruit, error) {
	if fruitEnt.Name == "" || fruitEnt.Price == 0{
		return fruitEntities.Fruit{}, constants.ErrEmptyInput
	}

	url := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%v", fruitEnt.Name)
	resp, err := http.Get(url)

    if err != nil {
		fmt.Println(err)
        return fruitEntities.Fruit{}, constants.ErrHitApi
    }
    defer resp.Body.Close()

	var fruitResp response.Fruit

	err = json.NewDecoder(resp.Body).Decode(&fruitResp)
	if err != nil {
		fmt.Println(err)
		return fruitEntities.Fruit{}, constants.ErrHitApi
	}

	temp := fruitEnt.Price

	fruitEnt = fruitEntities.Fruit{
		Name:  fruitResp.Name,
		Price: temp,
		Nutritions: fruitEntities.Nutrition {
			Calories: fruitResp.Nutritions.Calories,
			Fat: fruitResp.Nutritions.Fat,
			Sugar: fruitResp.Nutritions.Sugar,
			Carbohydrates: fruitResp.Nutritions.Carbohydrates,
			Protein: fruitResp.Nutritions.Protein,
		},
	}

	newUUID := uuid.New()
	uuidString := newUUID.String()
	fruitEnt.ID = uuidString

	fruitEnt, err = fruitUseCase.repository.AddFruit(fruitEnt)
	if err != nil {
		return fruitEntities.Fruit{}, constants.ErrInsertDatabase
	}

	return fruitEnt, nil
}

func (fruitUseCase *FruitUseCase) DeleteFruitById(id string) error {
	if id == "" {
		return constants.ErrEmptyInput
	}

	err := fruitUseCase.repository.DeleteFruitById(id)
	if err != nil {
		return constants.ErrGetDatabase
	}

	return nil
}