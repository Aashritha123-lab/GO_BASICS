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

├── main.go # Contains the HTTP server and handlers
└── README.md # Project documentation

## ▶️ How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/GO_BASICS.git
   cd GO_BASICS

2. Run the server:
bash
go run main.go
The server will start on: http://localhost:3051

API Endpoints

1. Create a Car
POST /cars
Request Body (JSON):
```bash
json
{
  "Name": "Civic",
  "Model": "2022",
  "Company": "Honda",
  "Year": 2022,
  "Price": 25000.50
}

Response:

json

{
  "ID": 123,
  "Name": "Civic",
  "Model": "2022",
  "Company": "Honda",
  "Year": 2022,
  "Price": 25000.5
}

2. Get a Car by ID
GET /cars/{id}

Example:

GET http://localhost:3051/cars/123

Response:

json

{
  "ID": 123,
  "Name": "Civic",
  "Model": "2022",
  "Company": "Honda",
  "Year": 2022,
  "Price": 25000.5
}

3. Delete a Car
DELETE /cars/{id}

Example:

DELETE http://localhost:3051/cars/123

Response:

Status: 200 OK

🛠️ Tech Stack
Language: Go

Packages: net/http, encoding/json, sync, math/rand

📌 Notes
Data is stored in memory only. If you restart the server, all cars are lost.

rand.Intn(1000) is used for generating car IDs → collisions are possible in rare cases.

For real-world use, you’d replace the in-memory map with a database (e.g., PostgreSQL, MongoDB).

🚀 Future Improvements
Add GET /cars to fetch all cars.

Add PUT /cars/{id} to update car details.

Replace random ID generation with proper UUIDs.

Persistent database storage.

👤 Author
Your Name
GitHub: @Aashritha-123-lab