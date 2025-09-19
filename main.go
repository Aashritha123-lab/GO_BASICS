package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Define the car struct for database
type Car struct {
	ID      int
	Name    string
	Model   string
	Company string
	Year    int
	Price   float64
}

// create the databse : "in memory maps"
var Cars = make(map[int]Car)

// to make concurrent operation seamlessly we use mutexes
var mu sync.Mutex

// create handler

// A handler is the code that handles an incoming HTTP request (reads it + sends a response).
func carhandler(w http.ResponseWriter, r *http.Request) {
	// &http.Request{
	//     Method: "GET",
	//     URL: &url.URL{
	//         Path: "/cars/8091",
	//     },
	//     Host: "localhost:3052",
	// }
	//s := r.URL.String() // Which will give entire string // suppose http request is like http://localhost:3052/cars/?id=8091 it will give you /cars/?id=8091
	url := r.URL.Path // ignoring the localhost will get the Path

	// suppose /cars/8091 it will trim to 8091
	entity := strings.TrimPrefix(url, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			CreateCar(w, r)
		} else {
			http.Error(w, "Invalid post request", http.StatusBadRequest) // where w http.ResponseWriter
		}
	case "GET":
		if entity == "" {
			http.Error(w, "Invalid GET Request", http.StatusNotFound)
		} else {
			id, _ := strconv.Atoi(entity)
			GetCar(w, r, id)
		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "Invalid DELETE Request", http.StatusNotFound)
		} else {
			id, _ := strconv.Atoi(entity)
			DeleteCar(w, r, id)
		}
	}
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Input Json format is wrong", http.StatusBadRequest)
	} // creating a new decoder which will read the JSON body of http request and then Deode function will write to the car struct

	id := rand.Intn(1000) // creating a random ID to save data in map
	car.ID = id           // updating the ID of car

	Cars[car.ID] = car
	w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)

	fmt.Println("Car has been created", http.StatusCreated)
}

func GetCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	car, ok := Cars[id]
	if !ok {
		http.Error(w, "Invalid Car ID", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(car)
	fmt.Println("Found the car:", http.StatusFound)
}

func DeleteCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := Cars[id]
	if !ok {
		http.Error(w, "Invalid Car ID", http.StatusNotFound)
		return
	}

	delete(Cars, id)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Deleted the car :", http.StatusOK)
}

// exposing to internet

func main() {
	http.HandleFunc("/cars", carhandler)
	http.HandleFunc("/cars/", carhandler)

	fmt.Println("HTTP server listening....")
	http.ListenAndServe(":3051", nil)
}
