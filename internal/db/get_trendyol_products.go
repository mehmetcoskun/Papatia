package db

import (
	"Papatia/internal/model"
	"log"
)

func GetTrendyolProducts() []model.Product {
	database := Connect()
	defer database.Close()

	rows, err := database.Query("select * from trendyol_products limit 200")
	if err != nil {
		log.Fatal(err)
	}

	var products []model.Product

	for rows.Next() {
		p := model.Product{}

		err := rows.Scan(&p.Id, &p.ProductCode, &p.ProductUrl, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, p)
	}

	defer rows.Close()
	return products
}
