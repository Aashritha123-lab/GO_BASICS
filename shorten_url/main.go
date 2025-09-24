package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type ShortenRequest struct {
	Url string
}

type Response struct {
	ResponseUrl    string
	ShortenRequest ShortenRequest
}

func generateCode() string {
	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)

	for i := range code {
		code[i] = letter[Rand.Intn(len(letter))]
	}

	return string(code)
}
func DBConnection() {
	dsn := "user=postgres password = Helper@123 dbname= postgres sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	DB = db
	fmt.Println("Connected to database successfully!")

}

func (url *Response) Insert() error {

	query := `INSERT INTO short_url(short_response,url)VALUES($1,$2)RETURNING short_response,url`

	if err := DB.QueryRow(query, url.ResponseUrl, url.ShortenRequest.Url).Scan(&url.ResponseUrl, &url.ShortenRequest.Url); err != nil {
		fmt.Printf("Error inserting to database :%v\n", err)
		return err
	}
	return nil
}

func (url *Response) Get() error {
	query := `SELECT short_response,url FROM short_url where short_response = $1`
	fmt.Println("url.ResponseUrl in get method is:", url.ResponseUrl)
	if err := DB.QueryRow(query, url.ResponseUrl).Scan(&url.ResponseUrl, &url.ShortenRequest.Url); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error in getting the url :%v\n", err)
			return err
		}
	}
	return nil
}

func (url *Response) UrlValidate() bool {
	// 1. Check if the original URL already exists
	var existingUrl string
	query := `SELECT url FROM short_url WHERE url = $1 LIMIT 1`
	err := DB.QueryRow(query, url.ShortenRequest.Url).Scan(&existingUrl)
	if err == nil {
		fmt.Println("URL already exists:", existingUrl)
		return true
	}
	if err != sql.ErrNoRows {
		fmt.Println("DB error while checking URL:", err)
		return true // fail safe
	}
	return false
}
func (url *Response) CodeValidate() bool {
	// 2. Check if the short code already exists
	var existingShort string
	query := `SELECT short_response FROM short_url WHERE short_response = $1 LIMIT 1`
	err := DB.QueryRow(query, url.ResponseUrl).Scan(&existingShort)
	if err == nil {
		fmt.Println("Short code already exists:", existingShort)
		return true
	}
	if err != sql.ErrNoRows {
		fmt.Println("DB error while checking short code:", err)
		return true
	}

	// No conflicts → safe to insert
	return false
}

func ShorthandUrl(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	short := &ShortenRequest{}
	url := &Response{
		ShortenRequest: *short,
	}
	if err := json.NewDecoder(r.Body).Decode(&short); err != nil {
		http.Error(w, "Invalid Json format", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(short.Url, "http://") && !strings.HasPrefix(short.Url, "https://") {
		short.Url = "https://" + short.Url
	}
	url.ShortenRequest.Url = short.Url

	var code string
	if url.UrlValidate() {
		http.Error(w, "url already exists", http.StatusConflict)
		return
	}

	for {
		code = generateCode()
		url.ResponseUrl = code
		if !url.CodeValidate() {

			break
		}
		fmt.Println("short code already exists, generating a new one...")
	}

	if err := url.Insert(); err != nil {
		http.Error(w, "Error inserting the url", http.StatusInternalServerError)
		return
	}

	resp := Response{
		ResponseUrl:    "http://localhost:3051/" + code,
		ShortenRequest: ShortenRequest{Url: short.Url},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func Redirect(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	requestedcode := strings.TrimPrefix(r.URL.Path, "/")

	url := &Response{}
	url.ResponseUrl = requestedcode

	if err := url.Get(); err != nil {
		http.Error(w, "Redirect Url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.ShortenRequest.Url, http.StatusFound)

}

var Rand *rand.Rand

func main() {
	DBConnection()
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	http.HandleFunc("/shorten", ShorthandUrl)
	http.HandleFunc("/", Redirect)

	fmt.Println("Http server is listening....")
	http.ListenAndServe(":3051", nil)

}
