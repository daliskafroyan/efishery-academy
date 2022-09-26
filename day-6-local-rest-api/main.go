package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	products = map[int]*Product{}
	nomor = 1
)

func createProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}

	if err := c.Bind(p); err != nil {
		return err
	}

	products[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, p)
}

func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, products[id])
}

func updateProduct(c echo.Context) error {
	p := new(Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	products[id].Name = p.Name
	return c.JSON(http.StatusOK, products[id])
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(products, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/products", createProduct)
	e.GET("/products", getAllProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}