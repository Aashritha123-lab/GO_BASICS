# 🚗 Go Cars REST API

This project is a simple **RESTful API in Go** that manages a collection of cars.  
It uses **in-memory storage** (`map[int]Car`) along with a **mutex** for thread-safe operations.  

The API supports:
- Creating a car (`POST /cars`)
- Fetching a car by ID (`GET /cars/{id}`)
- Deleting a car by ID (`DELETE /cars/{id}`)

---

## 📦 Features

- Written in **Go** using only the standard library (`net/http`, `encoding/json`).
- Simple **in-memory database** (no external dependencies).
- Thread-safe using **sync.Mutex**.
- JSON input/output for easy integration.

---

## 🏗️ Project Structure

.
├── main.go # Contains the HTTP server and handlers
└── README.md # Project documentation

yaml


---

## ▶️ How to Run

### 1. Clone the repository
	
	git clone https://github.com/<your-username>/GO_BASICS.git
	cd GO_BASICS 
	


### 2. Run the server

	 go run main.go



### 3. Access the server

	The server will start on:
	http://localhost:3051
	
	
## 🔗 API Endpoints
	### 1. Create a Car

		POST /cars

		Request Body (JSON):

		{
		  "Name": "Civic",
		  "Model": "2022",
		  "Company": "Honda",
		  "Year": 2022,
		  "Price": 25000.50
		}


		Response:

		{
		  "ID": 123,
		  "Name": "Civic",
		  "Model": "2022",
		  "Company": "Honda",
		  "Year": 2022,
		  "Price": 25000.5
		}