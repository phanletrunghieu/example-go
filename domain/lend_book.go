package domain

import (
	"time"
)

type LendBook struct {
	Model
	User_ID UUID      `sql:",type:uuid" json:"user_id"`
	Book_ID UUID      `sql:",type:uuid" json:"book_id"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}
