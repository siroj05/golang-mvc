package entities

import (
	"database/sql"
	"time"
)

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
