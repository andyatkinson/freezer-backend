package main

import (
  "os"
  "fmt"
  "time"
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

type Item struct {
  gorm.Model `json:"model"`
  Name string `gorm:"unique;not null" json:"name"`
  AddedOn string `json:"addedOn"`
}

func allItems(db *gorm.DB) func(echo.Context) error {
  return func(c echo.Context) error {
    var items []Item
    db.Find(&items)
    fmt.Println("{}", items)
    return c.JSON(http.StatusOK, items)
  }
}

func saveItem(db *gorm.DB) func(echo.Context) error {
  return func(c echo.Context) error {
    i := new(Item)
    // NOTE: default binder supports decoding application/json
    // binding is below
    if err = c.Bind(i); err != nil {
      return err
    }
    db.Create(&i)

    return c.JSON(http.StatusCreated, i)
  }
}

func main() {
  e := echo.New()

  db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Println(err.Error())
    panic("failed to connect database")
  }
  db.DB().SetConnMaxLifetime(time.Minute*5);
  db.DB().SetMaxIdleConns(0);
  db.DB().SetMaxOpenConns(5);
  //defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Item{})

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

  e.POST("/items", saveItem(db))
  e.GET("/items", allItems(db))

  e.Logger.Fatal(e.Start(":1323"))
}
