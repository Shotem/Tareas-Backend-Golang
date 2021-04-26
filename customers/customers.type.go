package customers

import "database/sql"

type Customer struct {
	CustomerID   string         `gorm:"primaryKey;column:CustomerID"`
	CompanyName  string         `gorm:"column:CompanyName"`
	ContactName  sql.NullString `gorm:"column:ContactName"`
	ContactTitle sql.NullString `gorm:"column:ContactTitle"`
	Country      sql.NullString `gorm:"column:Country"`
}

func (Customer) TableName() string { return "Customers" }
