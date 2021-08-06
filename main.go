package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type (
	employees struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	employee = map[int]*employees{}
	seq = 1
)

// create user
func createUser(c echo.Context) error {
	u := &employees{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	employee[u.ID] = u
	fmt.Println(employee[seq])
	seq++
	return c.JSON(http.StatusCreated, u)
}

// retrieve users
func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, employee)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, employee[id])
}

func updateUser(c echo.Context) error {
	u := new(employees)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	employee[id].Name = u.Name
	return c.JSON(http.StatusOK, employee[id])
}


func main() {
	e := echo.New()

	// logs the information about each HTTP request.
	e.Use(middleware.Logger())
	// recover from panic in  chain
	e.Use(middleware.Recover())

	//route
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)



	//start server
	e.Logger.Fatal(e.Start(":1323"))
}
