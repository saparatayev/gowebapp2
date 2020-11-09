package routes

import (
	"gowebapp2/models"
	"gowebapp2/utils"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	allProducts, allUsers, err := LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	utils.ExecuteTemplate(w, "admin.html", struct {
		AllProducts int64
		AllUsers    int64
	}{
		AllProducts: allProducts,
		AllUsers:    allUsers,
	})
}

func LoadData() (int64, int64, error) {
	allProducts, err := models.Count("products")
	if err != nil {
		return 0, 0, err
	}

	allUsers, err := models.Count("users")
	if err != nil {
		return 0, 0, err
	}

	return allProducts, allUsers, nil
}
