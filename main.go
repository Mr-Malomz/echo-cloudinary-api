package main

import (
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
          return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
    })

    e.Logger.Fatal(e.Start(":6000"))
}