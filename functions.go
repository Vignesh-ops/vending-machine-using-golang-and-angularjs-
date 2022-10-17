package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/rs/cors"
	//"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

func getquantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Product
	db := connect()
	result, err := db.Query("SELECT id,quantity from product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Product
		err := result.Scan(&post.ID, &post.Quantity)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	fmt.Print(posts)
	json.NewEncoder(w).Encode(posts)
}

func addproduct(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `INSERT INTO products(name,price)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyv := make(map[string]string)
	fmt.Println("value of body is", body)
	fmt.Printf("value of body is %s and type %T", body, body)
	fmt.Println("value of keyv is", keyv)
	json.Unmarshal(body, &keyv)
	//ID := keyv["id"]
	Name := keyv["name"]
	Price := keyv["price"]

	fmt.Print(Name, Price)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Price).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	var response = Product{}
	//var a interface{}
	fmt.Print(err)
	json.NewEncoder(w).Encode(response)
	defer db.Close()
}

func addproduct2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `INSERT INTO products(name,price)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyv := make(map[string]string)
	json.Unmarshal(body, &keyv)
	//ID := keyv["id"]
	Name := keyv["name"]
	Price := keyv["price"]
	fmt.Print(Name, Price)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Price).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	var response = Product{}
	//var a interface{}
	fmt.Print(err)
	json.NewEncoder(w).Encode(response)
	defer db.Close()
}
func addproduct3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `INSERT INTO products(name,price)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyv := make(map[string]string)
	json.Unmarshal(body, &keyv)
	//ID := keyv["id"]
	Name := keyv["name"]
	Price := keyv["price"]
	fmt.Print(Name, Price)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Price).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	var response = Product{}
	//var a interface{}
	fmt.Print(err)
	json.NewEncoder(w).Encode(response)
	defer db.Close()
}
func addproduct4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `INSERT INTO products(name,price)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyv := make(map[string]string)
	json.Unmarshal(body, &keyv)
	//ID := keyv["id"]
	Name := keyv["name"]
	Price := keyv["price"]
	fmt.Print(Name, Price)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Price).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	var response = Product{}
	//var a interface{}
	fmt.Print(err)
	json.NewEncoder(w).Encode(response)
	defer db.Close()
}
func addproduct5(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `INSERT INTO products(name,price)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyv := make(map[string]string)
	json.Unmarshal(body, &keyv)
	//ID := keyv["id"]
	Name := keyv["name"]
	Price := keyv["price"]
	fmt.Print(Name, Price)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Price).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	var response = Product{}
	//var a interface{}
	fmt.Print(err)
	json.NewEncoder(w).Encode(response)
	defer db.Close()
}

var err error

//updating product stocks to database
func updateqty(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db := connect()
	sqlStatement := `UPDATE product SET quantity = quantity-1 WHERE id = 1;`
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("updated")
	defer db.Close()
}

//get product by id
func getbyid(w http.ResponseWriter, r *http.Request) {
	db := connect()
	result, err := db.Query("SELECT quantity FROM product WHERE id=1")
	if err != nil {
		panic(err)
	}
	defer result.Close()
	var post Product
	for result.Next() {
		err := result.Scan(&post.Quantity)
		if err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(post)
}
