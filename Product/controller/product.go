package controller

import (
	"database/sql"
	"fmt"
	"product/model"
)

func InsertProduct(db *sql.DB, product model.Product) (msg string, err error) {
	query := ` CREATE TABLE IF NOT EXISTS product(
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price BIGINT NOT NULL,
			stock_quantity INT NOT NULL,
			description TEXT,
			added_on TIMESTAMP NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {

		return "", fmt.Errorf("error while creating table: %w", err)
	}
	query = `INSERT INTO product (id, name, price, stock_quantity, description, added_on)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = db.Exec(query, product.ID, product.Name, product.Price, product.StockQuantity, product.Description, product.AddedOn)
	if err != nil {
		err = fmt.Errorf("error while inserting product: %w", err)
		return "", err
	}
	msg = fmt.Sprintf("Product register with id : %v", product.ID)
	return msg, nil

}
