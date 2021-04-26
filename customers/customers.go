package customers

import (
	connect_utils "api/conn_utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCustomers(c echo.Context) error {
	var employees []Customer

	db := connect_utils.DB_info.Open()
	result := db.Find(&employees)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, employees)
}

func PostCustomer(c echo.Context) error {

	emp := Customer{
		CompanyName: c.FormValue("company-name"),
		ContactName: sql.NullString{String: c.FormValue("contact-Name"), Valid: true},
	}

	db := connect_utils.DB_info.Open()
	result := db.Select("LastName", "FirstName").Create(&emp)
	if result.Error != nil {
		log.Panicln("[Insert Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Succesfully")
}

func DeleteCustomer(c echo.Context) error {

	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("delete from %s where CustomerID = %d", Customer{}.TableName(), id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")

}

func PutCustomer(c echo.Context) error {
	title := c.FormValue("title")
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("update %s set CustomerTitle = '%s' where CustomerID = %d", Customer{}.TableName(), title, id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
