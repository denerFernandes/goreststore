package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/denerFernandes/goreststore/models"
	productRepository "github.com/denerFernandes/goreststore/repository/product"
	"github.com/denerFernandes/goreststore/utils"
	base64Upload "github.com/heliojuniorkroger/golang-base64-upload"
)

var products []models.Product

// AddProduct - insert product on database
func (c Controller) AddProduct(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		var error models.Error

		json.NewDecoder(r.Body).Decode(&product)

		if product.Title == "" {
			error.Message = "Title is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if product.Price == 0 {
			error.Message = "Price is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		imageName := utils.SanitizeString(product.Title) + ".png"

		err := base64Upload.Upload("upload/"+imageName, product.Image)
		if err != nil {
			panic(err)
		}

		product.Image = imageName

		productRepo := productRepository.ProductRepository{}
		product, err = productRepo.AddProduct(db, product)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error insert product"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			}
			utils.LogFatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, product)

	}

}
