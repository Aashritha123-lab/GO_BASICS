package handlers

import (
	"car/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// to make concurrent operation seamlessly we use mutexes
var Mu sync.Mutex

func Carhandler(w http.ResponseWriter, r *http.Request) {
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
	case "PUT":
		if entity == "" {
			http.Error(w, "Invalid DELETE Request", http.StatusNotFound)
		} else {
			id, _ := strconv.Atoi(entity)
			UpdateCar(w, r, id)
		}
	default:
		fmt.Println("Invalid method", http.StatusMethodNotAllowed)
	}
}
func CreateCar(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Input Json format is wrong", http.StatusBadRequest)
	} // creating a new decoder which will read the JSON body of http request and then Deode function will write to the car struct

	car.Insert()

	fmt.Println("Car saved to the inventory with id", car.ID)
	w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)

}

func GetCar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	car.ID = id
	err := car.Get()

	if err != nil {
		w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Error with no rows:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(car)
	fmt.Println("Found the car:", http.StatusFound)
}

func DeleteCar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	car.ID = id

	err := car.Delete()
	if err != nil {
		fmt.Printf("Error deleting the car with id: %v\n", car.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Deleted the car :", http.StatusOK)
}

func UpdateCar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Input Json format is wrong", http.StatusBadRequest)
	} // creating a new decoder which will read the JSON body of http request and then Deode function will write to the car struct

	car.ID = id

	err := car.Update()

	if err != nil {
		w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
		w.WriteHeader(http.StatusNotModified)
		return
	}

	fmt.Println("Car updated to the inventory with id", car.ID)

	w.Header().Set("Content-Type", "application/json") // it will tells the response body that Content-Type what type of data it is
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(car)
}
