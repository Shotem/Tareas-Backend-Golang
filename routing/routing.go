package routing

import (
	employees "api/employees"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	r := e.Group("/api")
	employeeRoutes(r)
}

func employeeRoutes(r *echo.Group) {
	r.GET("/employees", employees.GetEmployees)
	r.POST("/employees", employees.PostEmployee)
	r.DELETE("/employees", employees.DeleteEmployee)
	//r.DELETE("/cols", employees.GetCols)
}
