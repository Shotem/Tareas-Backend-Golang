package employees

import "database/sql"

type Employee struct {
	EmployeeID      int64          `gorm:"primaryKey;column:EmployeeID"`
	LastName        string         `gorm:"column:LastName"`
	FirstName       string         `gorm:"column:FirstName"`
	Title           sql.NullString `gorm:"column:Title"`
	TitleOfCourtesy sql.NullString `gorm:"column:TitleOfCourtesy"`
}

func (Employee) TableName() string { return "Employees" }
