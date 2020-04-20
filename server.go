package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type Item struct {
  Name string `json:"name" xml:"name" form:"name" query:"name"`
  AddedOn string `json:"addedOn" xml:"addedOn" form:"addedOn" query:"addedOn"`
}

func getItems(c echo.Context) error {
  var a1 = Item{Name: "meat", AddedOn: "2020-01-01"}
  var a2 = Item{Name: "veggies", AddedOn: "2020-02-01"}
  var a3 = Item{Name: "popsicles", AddedOn: "2020-03-01"}
  var a = []Item{a1, a2, a3}

  return c.JSON(http.StatusOK, a)
}

func saveItem(c echo.Context) error {
  i := new(Item)
  if err := c.Bind(i); err != nil {
    return err
  }

  // just return it as created
  return c.JSON(http.StatusCreated, i)
}

func main() {
  e := echo.New()

  // CORS default
  // Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
  // e.Use(middleware.CORS())

  // CORS restricted
  // Allows requests from any `http://localhost`
  // wth GET, PUT, POST or DELETE method.
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
  }))

  e.POST("/items", saveItem)
  e.GET("/items", getItems)

  e.Logger.Fatal(e.Start(":1323"))
}
