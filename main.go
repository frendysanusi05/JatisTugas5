package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type Customer struct {
    customer_id   string `json:"customerid"`
    name 		  string `json:"name"`
	address   	  string `json:"address"`
    email 	 	  string `json:"email"`
	phone_number  string `json:"phonenumber"`
}

type Driver struct {
    driver_id     string `json:"driverid"`
    name 		  string `json:"name"`
	phone_number  string `json:"phonenumber"`
}

type Order struct {
    order_id   					string `json:"orderid"`
    order_details 		  		string `json:"orderdetails"`
	location   	  				string `json:"location"`
    price_before_using_promo	string `json:"price_before_using_promo"`
}

type Promo struct {
    promo_id   		string `json:"promoid"`
    promo_details 	string `json:"promodetails"`
	valid_date   	string `json:"validdate"`
    discount 	 	string `json:"discount"`
}

type Gopay struct {
    payment_id  string `json:"paymentid"`
    balance 	string `json:"balance"`
}

type Cash struct {
    payment_id   string `json:"paymentid"`
}

type Make_order struct {
    customer_id   string `json:"customerid"`
    driver_id 	  string `json:"driverid"`
	order_id   	  string `json:"orderid"`
}

type Pay struct {
    customer_id   string `json:"customerid"`
    driver_id 	  string `json:"driverid"`
	order_id   	  string `json:"orderid"`
    payment_id 	  string `json:"paymentid"`
}

type Use_promo struct {
    order_id   string `json:"orderid"`
    promo_id   string `json:"promoid"`
}

type Payment struct {
    payment_id    string `json:"paymentid"`
    payment_type  string `json:"paymenttype"`
}

type Food struct {
	customer	Customer
	driver		Driver
	order		Order
	promo		Promo
	gopay		Gopay
	cash		Cash
	make_order	Make_order
	pay			Pay
	use_promo	Use_promo
	payment		Payment
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Food `json:"data"`
    Message string `json:"message"`
}

const (
    DB_USER     = "xxx" // insert your Postgre username
    DB_PASSWORD = "xxx" // insert your Postgre password
    DB_NAME     = "food_db"
)

func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return DB
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// response and request handlers
func GetDatabase(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    // Get all columns from food_db
    rows, err := db.Query("SELECT * FROM food_db")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var food []Food
	
    for rows.Next() {
        var id int
		var customer Customer
		var driver Driver
		var order Order
		var promo Promo
		var gopay Gopay
		var cash Cash
		var make_order Make_order
		var pay Pay
		var use_promo Use_promo
		var payment Payment

        err = rows.Scan(&id, &food)

        // check errors
        checkErr(err)

        food = append(movies, Food{Customer: customer, Driver: driver, Order: order, 
			Promo: promo, Gopay: gopay, Cash: cash, Make_order: make_order, 
			Pay: pay, Use_promo: use_promo, Payment: payment})
    }

    var response = JsonResponse{Type: "success", Data: food}

    json.NewEncoder(w).Encode(response)
}

func main() {
    router := mux.NewRouter()

// Route handles & endpoints
    // Get all database
    router.HandleFunc("/database/", GetDatabase).Methods("GET")
    fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8000", router))
}