package database

import (
	"alta/project/config"
	"alta/project/model"
)

func InputItem(products model.Product) (interface{}, error) {
	if err := config.DB.Save(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

//GET Product in shopping_cart
func GetProductInCart() (interface{}, error) {
	var shopping_carts []model.Shopping_cart
	if err := config.DB.Find(&shopping_carts).Error; err != nil {
		return nil, err
	}
	return shopping_carts, nil
}