package model

type Company struct {
	ID       int    `db:"id"`
	Code     string `db:"code"`
	Name     string `db:"name"`
	IsActive int8   `db:"is_active"`
}
