package contacts

import (
	"rva_crm/internal/core"
)

type Role string

const (
	RoleAdmin Role = "executive"
	RoleManager Role = "manager"
	RoleSales Role = "sales"
	RoleMisc Role = "misc"
)

type Contact struct {
	core.BaseModel

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	JobTitle  string `json:"job_title"`
}

type Address struct {
	Street string
	Street2 string
	City string
	State string
	Zip string
	Country string
}