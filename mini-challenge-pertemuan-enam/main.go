package main

import (
	"errors"
	"fmt"
	"mini-challenge-pertemuan-enam/database"
	"mini-challenge-pertemuan-enam/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createProduct("Biscuits")
	// updateProduct(1, "Soft Drink")
	// getProductById(1)
	// deleteProductById(1)
	// createVariant("Marie Regal", 38, 3)
	// updateVariantStock(2, 15)
	// deleteVariantById(2)
	getProductWithVariant()
}

// PRODUCT FUNCTION
func createProduct(name string) {
	db := database.GetDB()

	product := models.Product{
		Name: name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error creating product data: ", err)
		return
	}

	fmt.Println("New product data: ", product)
}

func updateProduct(id int, name string) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Model(&product).Where("id = ?", id).Updates(models.Product{Name: name}).Error

	if err != nil {
		fmt.Println("Error updating product data:", err)
		return
	}
	fmt.Printf("Update product's name: %+v \n", product.Name)
}

func getProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.First(&product, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			return
		}
		print("Error finding product:", err)
	}

	fmt.Printf("Product data: %+v \n", product)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}

func getProductWithVariant() {
	db := database.GetDB()

	products := []models.Product{}
	err := db.Preload("Variants").Find(&products).Error

	if err != nil {
		fmt.Println("Error getting product data with variants:", err.Error())
		return
	}

	fmt.Println("Product Datas With Variants")
	for _, product := range products {
		fmt.Println("Product:", product, "\n")
	}
}

// VARIANT FUNCTION
func createVariant(name string, stock int, productId uint) {
	db := database.GetDB()

	variant := models.Variant{
		VariantName: name,
		Stock:       stock,
		ProductID:   productId,
	}

	err := db.Create(&variant).Error

	if err != nil {
		fmt.Println("Error creating variant data: ", err)
		return
	}

	fmt.Println("New variant data: ", variant)
}

func updateVariantStock(id int, stock int) {
	db := database.GetDB()

	variant := models.Variant{}

	err := db.Model(&variant).Where("id = ?", id).Updates(models.Variant{Stock: stock}).Error

	if err != nil {
		fmt.Println("Error updating variant data:", err)
		return
	}
	fmt.Printf("Update variant's stock with id %v : %v", id, variant.Stock)
}

func deleteVariantById(id uint) {
	db := database.GetDB()

	variant := models.Variant{}

	err := db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		fmt.Println("Error deleting variant:", err.Error())
		return
	}

	fmt.Printf("Variant with id %d has been successfully deleted", id)
}
