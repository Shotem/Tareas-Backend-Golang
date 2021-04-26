package employees

import (
	connect_utils "api/conn_utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetEmployees(c echo.Context) error {
	var employees []Employee

	db := connect_utils.DB_info.Open()
	result := db.Find(&employees)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, employees)
}

func PostEmployee(c echo.Context) error {
	emp := Employee{
		LastName:  c.FormValue("lName"),
		FirstName: c.FormValue("fName"),
	}

	db := connect_utils.DB_info.Open()
	result := db.Select("LastName", "FirstName").Create(&emp)
	if result.Error != nil {
		log.Panicln("[Insert Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Succesfully")
}

// #region If GORM worked
/*
func DeleteEmployee(c echo.Context) error {

	db := connect_utils.DB_info.Open()

	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	result := db.Unscoped().Delete(&Employee{}, id)
	if result.Error != nil {
		log.Panicln("[Delete Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Deleted Succesfully")

}
*/
// #endregion

func DeleteEmployee(c echo.Context) error {

	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("delete from %s where EmployeeID = %d", Employee{}.TableName(), id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")

}

// #region If GORM worked
/*
func PutEmployee(c echo.Context) error {
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	var emp Employee

	db := connect_utils.DB_info.Open()
	result := db.First(&emp, id)
	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
	}

	emp.Title.String = c.FormValue("title")
	emp.Title.Valid = true

	result = db.Save(&emp)
	if result.Error != nil {
		log.Panicln("[Update Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
*/
// #endregion

func PutEmployee(c echo.Context) error {
	title := c.FormValue("title")
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	db, err := sql.Open("mssql", connect_utils.DB_info.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}

	result, err := db.Query(fmt.Sprintf("update %s set Title = '%s' where EmployeeID = %d", Employee{}.TableName(), title, id))
	if err != nil {
		log.Panic("[Update Error] " + err.Error())
	}
	defer result.Close()

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
