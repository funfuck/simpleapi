package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type product struct {
	ID          string
	Title       string
	Image       string
	Description string
}

// all product
var products = []product{
	{
		Title:       "apple",
		Image:       "http://icons.iconarchive.com/icons/bingxueling/fruit-vegetables/256/apple-red-icon.png",
		Description: "desc 1",
	},
	{
		Title:       "mangosteen",
		Image:       "https://cdn4.iconfinder.com/data/icons/Freshy/PNG/256/Mangosteen.png",
		Description: "desc 2",
	},
	{
		Title:       "watermelon",
		Image:       "http://icons.veryicon.com/ico/Leisure/Japan%20Summer/Watermelon%20cuts.ico",
		Description: "desc 3",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product", GetProduct)
	router.HandleFunc("/product/{id}", GetOneProduct)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// set json header and allow origin
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9999")

	// write response data
	json.NewEncoder(w).Encode(products)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {

	// get product id as index of product array
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	// set json header and allow origin
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9999")

	// get product and write response data
	product := products[idInt]
	json.NewEncoder(w).Encode(product)
}
