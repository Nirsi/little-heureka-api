package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var products []Product
var categories []Category
var offers []Offer

func InitializeStorage() {
	files, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".json") {
			content, err := os.ReadFile("./data/" + fileName)
			if err != nil {
				log.Fatal(err)
			}

			switch fileName {
			case "categories.json":
				err = json.Unmarshal(content, &categories)
			case "products.json":
				err = json.Unmarshal(content, &products)
			case "offers.json":
				err = json.Unmarshal(content, &offers)
			}

			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func ReadCategory(id int) Category {
	for _, category := range categories {
		if category.ID == id {
			return category
		}
	}
	return Category{}
}

func ReadCategories() []Category {
	return categories
}

func ReadProduct(id int) Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}

func ReadProducts(categoryId *int) []Product {
	var productsInCategory []Product
	for _, product := range products {
		if categoryId == nil || product.CategoryID == *categoryId {
			productsInCategory = append(productsInCategory, product)
		}
	}
	return productsInCategory
}

func ReadOffer(id int) Offer {
	for _, offer := range offers {
		if offer.ID == id {
			return offer
		}
	}
	return Offer{}
}

func ReadOffers(productId *int) []Offer {
	var offersInCategory []Offer
	for _, offer := range offers {
		if productId == nil || offer.ProductID == *productId {
			offersInCategory = append(offersInCategory, offer)
		}
	}
	return offersInCategory
}
