package model

import "time"

type Auth struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
