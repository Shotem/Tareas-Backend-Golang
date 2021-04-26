package main

import (
	connect_utils "api/conn_utils"
	routing "api/routing"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {

	e := echo.New()
	connect_utils.DB_info = connect_utils.DB_Information{
		Server:   "localhost",
		Database: "Northwind",
		User:     "",
		Password: "",
		Port:     1433,
	}

	routing.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
