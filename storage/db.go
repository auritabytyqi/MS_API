package storage

import (
	"log"

	"MS_API/FOODS_MS/model"
	config "MS_API/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func AddFoodRecord(food model.Food) error {
	DB.Select("FoodId", "FoodName", "FoodDescription").Create(&food)
	return nil
}

func DeleteFood(id string) error {
	DB.Delete(&model.Food{}, id)
	return nil
}

func UpdateFood(id, name, description string) error {
	DB.Model(&model.Food{}).Where("id=?", id).Updates(model.Food{FoodName: name, FoodDescription: description})
	return nil
}
