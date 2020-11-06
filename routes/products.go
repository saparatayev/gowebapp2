package routes

import (
	"gowebapp2/models"
	"gowebapp2/sessions"
	"gowebapp2/utils"
	"net/http"
	"strconv"
)

func productsGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	total := uint64(len(products))

	message, alert := sessions.Flash(w, r)

	utils.ExecuteTemplate(w, "products.html", struct {
		Total    uint64
		Products []models.Product
		Alert    utils.Alert
	}{
		Total:    total,
		Products: products,
		Alert:    utils.NewAlert(message, alert),
	})
}

func productsCreateGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	utils.ExecuteTemplate(w, "product_create.html", struct {
		Categories []models.Category
	}{
		Categories: categories,
	})
}

func productsCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := verifyInputProduct(r)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	_, err = models.NewProduct(product)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	sessions.Message("New product is added", "success", w, r)

	http.Redirect(w, r, "/products", 302)
}

func verifyInputProduct(r *http.Request) (models.Product, error) {
	r.ParseForm()

	var product models.Product

	var err error = nil

	product.Name = r.PostForm.Get("name")

	product.Price, err = strconv.ParseFloat(r.PostForm.Get("price"), 64)
	if err != nil {
		return models.Product{}, err
	}

	product.Quantity, err = strconv.Atoi(r.PostForm.Get("quantity"))
	if err != nil {
		return models.Product{}, err
	}

	product.Amount = product.Price * float64(product.Quantity)

	product.Category.Id, _ = strconv.Atoi(r.PostForm.Get("category"))

	return product, nil
}

func productEditGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	productId, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)
}
