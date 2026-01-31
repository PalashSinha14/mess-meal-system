package models

import (
	"github.com/PalashSinha14/mess-meal-system/db"
)

type Meal struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	UserID      int64
}

var meals = []Meal{} 


func (m *Meal) Save() error { 
	query := `INSERT INTO meals(name, description, user_id) VALUES (?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.Name, m.Description, m.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	m.ID = id
	return err
}

func getAllMeals() ([]Meal, error) { 
	query := "SELECT * FROM meals"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var meals []Meal
	for rows.Next() {
		var meal Meal
		err := rows.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.UserID)
		if err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}
	return meals, nil
}

func getMealByID(id int64) (*Meal, error) {
	query := "SELECT * FROM meals WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var meal Meal
	err := row.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.UserID)
	if err != nil {
		return nil, err
	}
	return &meal, nil
}
