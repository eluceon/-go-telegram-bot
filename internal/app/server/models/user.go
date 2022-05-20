package models

import "database/sql"

type User struct {
	ID             int64
	Username       string
	CorrectAnswers int
	TotalAnswers   int
	IsPassing      bool
	RegisteredAt   sql.NullTime
}
