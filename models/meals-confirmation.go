package models

import "github.com/PalashSinha14/mess-meal-system/db"

type MealConfirmation struct {
	UserID int64  `json:"user_id"`
	MealID int64  `json:"meal_id"`
	Status string `json:"status"` // YES / NO
}

func (mc MealConfirmation) Save() error {
	_, err := db.DB.Exec(`
	INSERT INTO meal_confirmations(user_id, meal_id, status)
	VALUES (?,?,?)
	ON CONFLICT(user_id, meal_id)
	DO UPDATE SET status = excluded.status`,
		mc.UserID, mc.MealID, mc.Status,
	)
	return err
}

func GetUserConfirmation(userId int64) ([]MealConfirmation, error) {
	rows, err := db.DB.Query(
		"SELECT meal_id, status FROM meal_confirmations WHERE user_id = ?",
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []MealConfirmation
	for rows.Next() {
		var mc MealConfirmation
		mc.UserID = userId
		rows.Scan(&mc.MealID, &mc.Status)
		result = append(result, mc)
	}
	return result, nil
}
