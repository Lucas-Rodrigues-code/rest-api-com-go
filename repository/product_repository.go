package repository

import (
	"api-go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	var products_list []model.Product
	var product_obj model.Product

	for rows.Next() {
		err := rows.Scan(
			&product_obj.Id,
			&product_obj.Name,
			&product_obj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		products_list = append(products_list, product_obj)
	}
	rows.Close()

	return products_list, nil
}
