package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, Tmp{ID: 1, Name: "tmp"})
}

type Tmp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
