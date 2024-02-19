package productcontroller

import (
	"encoding/json"
	"fmt"
	"golang-native/entities"
	"golang-native/models/categorymodel"
	"golang-native/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	categories := categorymodel.GetAll()
	data := map[string]any{
		"products":   products,
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/products/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var product entities.Product

	// convert data
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		panic(err)
	}

	stock, err := strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		panic(err)
	}

	product.Name = r.FormValue("name")
	product.Category.Id = uint(categoryId)
	product.Price = int(price)
	product.Stock = int(stock)
	product.Desc = r.FormValue("desc")
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if ok := productmodel.Create(product); !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(r.Body)

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func Show(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	product := productmodel.GetById(id)

	jsonInBytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	var product entities.Product

	// convert data
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		panic(err)
	}

	stock, err := strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		panic(err)
	}

	product.Name = r.FormValue("name")
	product.Category.Id = uint(categoryId)
	product.Price = int(price)
	product.Stock = int(stock)
	product.Desc = r.FormValue("desc")
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if ok := productmodel.Update(id, product); !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	if ok := productmodel.Delete(id); !ok {
		temp, _ := template.ParseFiles("views/products/index.html")
		temp.Execute(w, nil)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
	})
}
