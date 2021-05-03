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
	r.GET("/employees/:id", employees.GetEmployeeByID)
	r.POST("/employees", employees.PostEmployee)
	r.DELETE("/employees/:id", employees.DeleteEmployee)
	r.PUT("/employees/:id", employees.PutEmployee)
}

func productRoutes(r *echo.Group) {
	r.GET("/products", product.GetProducts)
	r.GET("/products/:id", product.GetProductByID)
	r.POST("/products", product.PostProduct)
	r.DELETE("/products/:id", product.DeleteProduct)
	r.PUT("/products/:id", product.PutProduct)
}

func customerRoutes(r *echo.Group) {
	r.GET("/customers", customer.GetCustomers)
	r.GET("/customers/:id", customer.GetCustomerByID)
	r.POST("/customers", customer.PostCustomer)
	r.DELETE("/customers/:id", customer.DeleteCustomer)
	r.PUT("/customers/:id", customer.PutCustomer)
}
