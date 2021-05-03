package customers

import (
	connect_utils "api/conn_utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCustomers(c echo.Context) error {
	var customers []Customer

	db := connect_utils.DB_info.Open()
	result := db.Find(&customers)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, customers)
}

func GetCustomerByID(c echo.Context) error {
	id := c.Param("id")

	var customer Customer
	db := connect_utils.DB_info.Open()
	result := db.Where("CustomerID = ?", id).Find(&customer)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

func PostCustomer(c echo.Context) error {

	cus := Customer{
		CustomerID:  c.FormValue("id"),
		CompanyName: c.FormValue("company-name"),
		ContactName: sql.NullString{String: c.FormValue("contact-name"), Valid: true},
	}

	db := connect_utils.DB_info.Open()
	result := db.Select("CustomerID", "CompanyName", "ContactName").Create(&cus)
	if result.Error != nil {
		log.Panicln("[Insert Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Succesfully")
}

func DeleteCustomer(c echo.Context) error {

	id := c.Param("id")
	db := connect_utils.DB_info.Open()

	result := db.Delete(&Customer{}, "CustomerID = ?", id)
	if result.Error != nil {
		log.Panicln("[Delete Error] " + result.Error.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")
}

func PutCustomer(c echo.Context) error {
	title := c.FormValue("title")
	id := c.FormValue("id")

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("update %s set ContactTitle = '%s' where CustomerID = '%s'", Customer{}.TableName(), title, id))
	if err != nil {
		log.Panicln("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
