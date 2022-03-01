package entity

import (
	sqlx "github.com/mytoko2796/sdk-go/stdlib/sql"
)

type CreateTodoList struct {
	ID string `json:"id"`
	UserID int64 `json:"user_id"`
	Name string `json:"name"`
	StartOn sqlx.NullTime `json:"start_on"`
	Status int64 `json:"status"`
}
