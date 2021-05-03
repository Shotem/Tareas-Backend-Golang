package products

import (
	connect_utils "api/conn_utils"
	"database/sql"
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

func GetProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	var employees []Product

	db := connect_utils.DB_info.Open()
	result := db.Where("ProductID = ?", id).Find(&employees)

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
	result := db.Select("ProductName", "UnitPrice").Create(&prod)
	if result.Error != nil {
		log.Panicln("[Insert Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Succesfully")
}

func DeleteProduct(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	db := connect_utils.DB_info.Open()

	result := db.Delete(&Product{}, "ProductID = ?", id)
	if result.Error != nil {
		log.Panicln("[Delete Error] " + result.Error.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")

}

func PutProduct(c echo.Context) error {
	uPrice, err := strconv.ParseFloat(c.FormValue("uPrice"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	db := connect_utils.DB_info.Open()
	var product Product

	result := db.Find(&product, "ProductID = ?", id)
	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
	}
	defer result.Close()

	product.UnitPrice.Float64 = uPrice
	product.UnitPrice.Valid = true

	result = db.Model(&product).Update("UnitPrice", uPrice)
	if result.Error != nil {
		log.Panicln("[Update Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
