package models

import (
	"liquid/pkg/db"
)

func GenerateTigerId() int {
	res := db.QueryVal[int]("SELECT * FROM generate_tiger_id()")
	return res
}
