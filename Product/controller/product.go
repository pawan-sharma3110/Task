package controller

import (
	"database/sql"
	"fmt"
	"product/model"
)

func InsertProduct(db *sql.DB, product model.Product) (msg string, err error) {
	query := `
  id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL, -- Using BIGINT for large prices; adjust as needed
    stock_quantity INT NOT NULL, -- Use INT or BIGINT based on the expected range
    description TEXT, -- Changed "discription" to "description" for correct spelling
    added_on TIMESTAMP NOT NULL`
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
