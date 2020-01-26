package productrepository

import (
	"database/sql"

	"github.com/denerFernandes/goreststore/models"
	"github.com/denerFernandes/goreststore/utils"
)

// AddProduct - insert product on database
func (u ProductRepository) AddProduct(db *sql.DB, product models.Product) (models.Product, error) {

	sqlStmt := "insert into products (title, description, image, price, status) values($1, $2, $3, $4, $5) RETURNING id;"
	err := db.QueryRow(sqlStmt, product.Title, product.Description, product.Image, product.Price, product.Status).Scan(&product.ID)
	utils.LogFatal(err)

	return product, err
}
