package main

import (
	"fmt"
	"golang-native/config"
	"golang-native/controllers/categorycontroller"
	"golang-native/controllers/homecontroller"
	"golang-native/controllers/productcontroller"
	"log"
	"net/http"
	"os"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	// categories handler
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/show", categorycontroller.Show)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// products handler
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/show", productcontroller.Show)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		directoryPath := "/assets/products"

		// Memeriksa apakah direktori ada
		fileInfo, err := os.Stat(directoryPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Direktori tidak ditemukan")
			} else {
				fmt.Println("Gagal memeriksa direktori:", err)
			}
			return
		}

		if fileInfo.IsDir() {
			fmt.Println("Direktori ditemukan")
		} else {
			fmt.Println("Path tersebut bukan sebuah direktori")
		}
	})

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
