package models

import (
	"car/config"
	"database/sql"
	"errors"
	"fmt"
)

// var Cars = make(map[int]models.Car)

type Car struct {
	ID    int
	Name  string
	Model string
	Brand string
	Year  int
	Price float64
}

// INSERT INTO cars(name,model,brand,year,price)VALUES('X7','V8','BMW',2023,20000000);
func (c *Car) Insert() {
	query := `INSERT INTO cars(name,model,brand,year,price)VALUES($1,$2,$3,$4,$5) RETURNING id`
	if err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.ID); err != nil {
		fmt.Printf("Error inserting to database :%v\n", err)
	}
}

// GETTING THE DETAILS

func (c *Car) Get() error {
	query := `SELECT name, model,brand,year,price FROM cars where id = $1`
	if err := config.DB.QueryRow(query, c.ID).Scan(&c.Name, &c.Model, &c.Brand, &c.Year, &c.Price); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error inserting to database :%v\n", err)
			return err
		}
	}
	return nil
}

func (c *Car) Delete() error {
	query := `DELETE FROM cars where ID = $1`
	res, err := config.DB.Exec(query, c.ID)
	if row, _ := res.RowsAffected(); row == 0 && err == nil {
		return errors.New("Car with id not found")
	}
	return nil
}

// UPDATE cars SET name=$1,model=$2,brand=$3,year=$4,price=$5 WHERE id = ID

func (c *Car) Update() error {

	query := `UPDATE cars SET name=$1,model=$2,brand=$3,year=$4,price=$5 WHERE id = $6;`
	res, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.ID)
	if row, _ := res.RowsAffected(); row == 0 && err == nil {
		return errors.New("Car with id not found so nothing updated")
	}
	return nil
}
