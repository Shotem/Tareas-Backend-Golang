package employees

import "database/sql"

type Employee struct {
	EmployeeID      int64          `gorm:"primaryKey;column:EmployeeID"`
	LastName        string         `gorm:"column:LastName"`
	FirstName       string         `gorm:"column:FirstName"`
	Title           sql.NullString `gorm:"column:Title"`
	TitleOfCourtesy sql.NullString `gorm:"column:TitleOfCourtesy"`
	/*
		BirthDate       sql.NullTime   `gorm:"column:BirthDate"`
		HireDate        sql.NullTime   `gorm:"column:HireDate"`
		Address         sql.NullString `gorm:"column:Address"`
		City            sql.NullString `gorm:"column:City"`
		Region          sql.NullString `gorm:"column:Region"`
		PostalCode      sql.NullString `gorm:"column:PostalCode"`
		Country         sql.NullString `gorm:"column:Country"`
		HomePhone       sql.NullString `gorm:"column:HomePhone"`
		Extension       sql.NullString `gorm:"column:Extension"`
		Notes           sql.NullString `gorm:"column:Notes"`
		ReportsTo       sql.NullInt64  `gorm:"column:ReportsTo"`
		PhotoPath       sql.NullString `gorm:"column:PhotoPath"`*/
}

func (Employee) TableName() string { return "Employees" }
