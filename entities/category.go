package entities

import (
	"database/sql"
)

type Category struct {
	Id        uint
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
