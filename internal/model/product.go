package model

import "database/sql"

type Product struct {
	Id          int
	ProductCode string
	ProductUrl  string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   sql.NullString
}
