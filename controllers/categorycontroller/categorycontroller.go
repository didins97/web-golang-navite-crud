package categorycontroller

import (
	"encoding/json"
	"golang-native/entities"
	"golang-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Show(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	category := categorymodel.GetById(id)

	jsonInBytes, err := json.Marshal(category)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	category.Name = r.FormValue("name")
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	ok := categorymodel.Create(category)
	if !ok {
		temp, _ := template.ParseFiles("views/categories/index.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	var category entities.Category
	category.Name = r.FormValue("name")
	category.UpdatedAt = time.Now()

	ok := categorymodel.Update(category, id)
	if !ok {
		temp, _ := template.ParseFiles("views/categories/index.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	ok := categorymodel.Delete(id)
	if !ok {
		temp, _ := template.ParseFiles("views/categories/index.html")
		temp.Execute(w, nil)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
	})
}
