package routing

import (
	customer "api/customers"
	employees "api/employees"
	product "api/products"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	r := e.Group("/api")
	employeeRoutes(r)
	productRoutes(r)
	customerRoutes(r)
}

func employeeRoutes(r *echo.Group) {
	r.GET("/employees", employees.GetEmployees)
	r.POST("/employees", employees.PostEmployee)
	r.DELETE("/employees", employees.DeleteEmployee)
	r.PUT("/employees", employees.PutEmployee)
}

func productRoutes(r *echo.Group) {
	r.GET("/products", product.GetProducts)
	r.POST("/products", product.PostProduct)
	r.DELETE("/products", product.DeleteProduct)
	r.PUT("/products", product.PutProduct)
}

func customerRoutes(r *echo.Group) {
	r.GET("/customers", customer.GetCustomers)
	r.POST("/customers", customer.PostCustomer)
	r.DELETE("/customers", customer.DeleteCustomer)
	r.PUT("/customers", customer.PutCustomer)
}
