package products

import "database/sql"

type Product struct {
	ProductID       int64           `gorm:"primaryKey;column:ProductID"`
	ProductName     string          `gorm:"column:ProductName"`
	QuantityPerUnit string          `gorm:"column:QuantityPerUnit"`
	UnitPrice       sql.NullFloat64 `gorm:"column:UnitPrice"`
	UnitsInStock    sql.NullInt32   `gorm:"column:UnitsInStock"`
}

func (Product) TableName() string { return "Products" }
