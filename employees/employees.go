package employees

import (
	connect_utils "api/conn_utils"
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

func GetEmployeeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}

	var employee Employee
	db := connect_utils.DB_info.Open()
	result := db.Where("EmployeeID = ?", id).Find(&employee)

	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, employee)
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

func DeleteEmployee(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	db := connect_utils.DB_info.Open()

	result := db.Delete(&Employee{}, "EmployeeID = ?", id)
	if result.Error != nil {
		log.Panicln("[Delete Error] " + result.Error.Error())
	}
	defer result.Close()

	return c.String(http.StatusOK, "Deleted Succesfully")

}

func PutEmployee(c echo.Context) error {
	title := c.FormValue("title")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse "+c.FormValue("id"))
	}
	db := connect_utils.DB_info.Open()
	var employee Employee

	result := db.Find(&employee, "EmployeeID = ?", id)
	if result.Error != nil {
		log.Panicln("[Select Error] " + result.Error.Error())
	}
	defer result.Close()

	result = db.Model(&employee).Update("Title", title)
	if result.Error != nil {
		log.Panicln("[Update Error] " + result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Updated Succesfully")
}
