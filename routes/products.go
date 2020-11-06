package routes

import (
	"gowebapp2/models"
	"gowebapp2/utils"
	"net/http"
)

func productsGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	total := uint64(len(products))

	utils.ExecuteTemplate(w, "products.html", struct {
		Total    uint64
		Products []models.Product
	}{
		Total:    total,
		Products: products,
	})
}
