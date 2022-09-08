package controller

import (
	"net/http"

	"MS_API/RESTAURANTS_MS/model"
	"MS_API/storage"

	"github.com/labstack/echo/v4"
)

// func GetRestaurant(c echo.Context) error {
// 	id := c.Param("id")
// 	foods, _ := GetRepoFoods()
// 	for i := 0; i < len(foods); i++ {
// 		if foods[i].FoodId == id {
// 			return c.JSON(http.StatusOK, foods[i])
// 		}
// 	}
// 	return c.JSON(http.StatusOK, "Food doesn't exist")
// }

func GetRestaurants(c echo.Context) error {
	restaurants, _ := GetRepoRestaurants()
	return c.JSON(http.StatusOK, restaurants)
}

func GetRepoRestaurants() ([]model.Restaurant, error) {
	db := storage.GetDBInstance()
	var restaurants []model.Restaurant
	if err := db.Raw("SELECT * FROM Restaurants ").Scan(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// func AddFood(c echo.Context) error {
// 	id := c.QueryParam("id")
// 	name := c.QueryParam("name")
// 	description := c.QueryParam("description")
// 	id_exists := FoodExists(id)
// 	if !id_exists {
// 		food := model.Food{FoodId: id, FoodName: name, FoodDescription: description}
// 		storage.AddFoodRecord(food)
// 		return c.JSON(http.StatusOK, "Food is created")
// 	}
// 	return c.JSON(http.StatusOK, "Food product exists with this id. Try another one...")
// }

// func DeleteFood(c echo.Context) error {
// 	id := c.Param(("id"))
// 	id_exists := FoodExists(id)
// 	if id_exists {
// 		storage.DeleteFood(id)
// 		return c.JSON(http.StatusOK, "Food is deleted")
// 	}
// 	return c.JSON(http.StatusOK, "Food doesn't exist")
// }

// func UpdateFood(c echo.Context) error {
// 	id := c.QueryParam("id")
// 	name := c.QueryParam("name")
// 	description := c.QueryParam("description")
// 	id_exists := FoodExists(id)
// 	if id_exists {
// 		storage.UpdateFood(id, name, description)
// 		return c.JSON(http.StatusOK, "Food is updated")
// 	}
// 	return c.JSON(http.StatusOK, "Food doesn't exist")
// }

// func FoodExists(id string) bool {
// 	foods, _ := GetRepoFoods()
// 	id_exists := false
// 	for i := 0; i < len(foods); i++ {
// 		if foods[i].FoodId == id {
// 			id_exists = true
// 		}
// 	}
// 	return id_exists
// }
