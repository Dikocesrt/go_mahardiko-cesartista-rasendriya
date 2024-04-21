package fruit

import (
	fruitEntities "clean2/entities/fruit"

	"gorm.io/gorm"
)

type FruitRepo struct {
	DB *gorm.DB
}

func NewFruitRepo(db *gorm.DB) *FruitRepo {
	return &FruitRepo{
		DB: db,
	}
}

func (fruitRepo *FruitRepo) GetAllFruits() ([]fruitEntities.Fruit, error) {
	var fruitsDB []Fruit

	err := fruitRepo.DB.Find(&fruitsDB).Error
	if err != nil {
		return []fruitEntities.Fruit{}, err
	}

	var fruitsEntities []fruitEntities.Fruit

	for _, fruitDB := range fruitsDB {
		temp := fruitDB.fromFruitDBToFruitEntities()
		fruitsEntities = append(fruitsEntities, *temp)
	}

	return fruitsEntities, nil
}

func (fruitRepo *FruitRepo) GetFruitById(id string) (fruitEntities.Fruit, error) {
	var fruitDB Fruit
	err := fruitRepo.DB.Where("id = ?", id).First(&fruitDB).Error
	if err != nil {
		return fruitEntities.Fruit{}, err
	}

	fruitEnt := *fruitDB.fromFruitDBToFruitEntities()
	return fruitEnt, nil
}

func (fruitRepo *FruitRepo) AddFruit(fruitEnt fruitEntities.Fruit) (fruitEntities.Fruit, error) {
	fruitDB := fromFruitEntitiesToFruitDB(fruitEnt)

	err := fruitRepo.DB.Create(&fruitDB).Error
	if err != nil {
		return fruitEntities.Fruit{}, err
	}

	newFruit := *fruitDB.fromFruitDBToFruitEntities()

	return newFruit, nil
}

func (fruitRepo *FruitRepo) DeleteFruitById(id string) error {
	err := fruitRepo.DB.Where("id = ?", id).Delete(&Fruit{}).Error
	if err != nil {
		return err
	}
	return nil
}