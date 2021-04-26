package products

import (
	connect_utils "api/conn_utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	var employees []Product

	db := connect_utils.DB_info.Open()
	result := db.Find(&employees)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, employees)
}

func PostProduct(c echo.Context) error {
	uPrice, err := strconv.ParseFloat(c.FormValue("uPrice"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	prod := Product{
		ProductName: c.FormValue("pName"),
		UnitPrice:   sql.NullFloat64{Float64: uPrice, Valid: true},
	}

	db := connect_utils.DB_info.Open()
	result := db.Select("LastName", "FirstName").Create(&prod)
	if result.Error != nil {
		log.Panicln("[Insert Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Succesfully")
}

func DeleteProduct(c echo.Context) error {

	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("delete from %s where ProductID = %d", Product{}.TableName(), id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")

}

func PutProduct(c echo.Context) error {
	uPrice, err := strconv.ParseFloat(c.FormValue("uPrice"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("update %s set UnitPrice = %f where ProductID = %d", Product{}.TableName(), uPrice, id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
