package routes

import (
	"errors"
	"fmt"
	"gowebapp2/models"
	"gowebapp2/sessions"
	"gowebapp2/utils"
	"html"
	"net/http"
	"strconv"
)

var (
	ErrPriceValue          = errors.New("Price is not valid")
	ErrQuantityValue       = errors.New("Quantity is not valid")
	ErrRequiredProductName = errors.New("Product name is required")
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

	message, alert := sessions.Flash(w, r)

	utils.ExecuteTemplate(w, "product_create.html", struct {
		Categories []models.Category
		Alert      utils.Alert
	}{
		Categories: categories,
		Alert:      utils.NewAlert(message, alert),
	})
}

func productsCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := verifyInputProduct(r)
	if err != nil {
		sessions.Message(fmt.Sprintf("%s", err), "danger", w, r)

		http.Redirect(w, r, "/products/create", 302)

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

	product.Id, _ = strconv.ParseUint(r.PostForm.Get("id"), 10, 64)

	product.Name = html.EscapeString(r.PostForm.Get("name"))

	if models.IsEmpty(product.Name) {
		return product, ErrRequiredProductName
	}

	if !models.Max(product.Name, 255) {
		return product, models.ErrMaxLimit
	}

	product.Price, err = strconv.ParseFloat(r.PostForm.Get("price"), 64)
	if err != nil {
		return product, ErrPriceValue
	}

	product.Quantity, err = strconv.Atoi(r.PostForm.Get("quantity"))
	if err != nil {
		return product, ErrQuantityValue
	}

	product.Amount = product.Price * float64(product.Quantity)

	product.Category.Id, _ = strconv.Atoi(r.PostForm.Get("category"))

	return product, nil
}

func productEditGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	productId, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)

	product, err := models.GetProductById(productId)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	categories, err := models.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	priceFormat := product.PriceToString()

	message, alert := sessions.Flash(w, r)

	utils.ExecuteTemplate(w, "product_edit.html", struct {
		Categories  []models.Category
		Product     models.Product
		PriceFormat string
		Alert       utils.Alert
	}{
		Categories:  categories,
		Product:     product,
		PriceFormat: priceFormat,
		Alert:       utils.NewAlert(message, alert),
	})
}

func productEditPostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := verifyInputProduct(r)

	if err != nil {
		sessions.Message(fmt.Sprintf("%s", err), "danger", w, r)

		http.Redirect(w, r, fmt.Sprintf("/product/edit?productId=%d", product.Id), 302)

		return
	}

	rows, err := models.UpdateProduct(product)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	sessions.Message(fmt.Sprintf("Products %d info changed", rows), "info", w, r)

	http.Redirect(w, r, "/products", 302)
}

func productDeleteGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	id, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)

	ok, _ := strconv.ParseBool(keys.Get("confirm"))
	if !ok {
		http.Redirect(w, r, "/products", 302)
		return
	}

	rows, err := models.DeleteProduct(id)

	if err != nil {
		utils.InternalServerError(w)
		return
	}

	sessions.Message(fmt.Sprintf("Product %d deleted", rows), "warning", w, r)

	http.Redirect(w, r, "/products", 302)
}
