package routes

import (
	"gowebapp2/models"
	"gowebapp2/utils"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	products, users, err := LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	allProducts := int64(len(products))
	allUsers := int64(len(users))
	lastUser := users[len(users)-1]
	lastProduct := products[len(products)-1]

	utils.ExecuteTemplate(w, "admin.html", struct {
		AllProducts int64
		AllUsers    int64
		LastProduct models.Product
		LastUser    models.User
	}{
		AllProducts: allProducts,
		AllUsers:    allUsers,
		LastProduct: lastProduct,
		LastUser:    lastUser,
	})
}

func LoadData() ([]models.Product, []models.User, error) {
	products, err := models.GetProducts()
	if err != nil {
		return nil, nil, err
	}

	users, err := models.GetUsers()
	if err != nil {
		return nil, nil, err
	}

	return products, users, nil
}
