package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	Todo struct {
		ID     int    `json:"id"`
		Item   string `json:"item"`
		IsDone bool   `json:"isDone"`
	}

	TodoList struct {
		Title string
		Todos []Todo
	}
)

var (
	indexTemplate *template.Template
	todoList      *TodoList
	idNumber      = 3
)

//****************
// Route Handlers
//****************

// useful when using Docker or Kubernetes to see if the application is up
func getStatus(c echo.Context) error {
	c.Logger().Info("GET for status from " + c.Request().RemoteAddr)

	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

// used to get a webpage that shows all the todos
func getIndex(c echo.Context) error {
	c.Logger().Debug("GET for index from " + c.Request().RemoteAddr)

	// we will run our template on a buffer
	// this lets us check for errors
	var buffer bytes.Buffer
	err := indexTemplate.Execute(&buffer, todoList)
	if err != nil {
		// we log the error on server side,
		// but don't give any details to the client
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "There was a problem on the server")
	}
	//everything was ok with the template, so write it to the client
	return c.HTML(http.StatusOK, buffer.String())
}

// used to get a JSON representation of the Todo List
// Sample: /todos
func getTodos(c echo.Context) error {
	c.Logger().Debug("GET for todos from " + c.Request().RemoteAddr)

	return c.JSON(http.StatusOK, todoList)
}

// used to get a specific Todo based on ID
// Sample: /todo/1
func getTodo(c echo.Context) error {
	// todo ID from path `todo/:id`
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error("Get for id failed on " + c.Param("id"))
		return c.HTML(http.StatusBadRequest, "Get for id failed on "+c.Param("id"))
	}

	c.Logger().Debug(fmt.Sprintf("GET for todo from %s for id:%d", c.Request().RemoteAddr, id))

	if id > 2 || id <= 0 {
		// for the sample we only consider 2 items
		c.Logger().Error(fmt.Sprintf("%d not found", id))
		return c.HTML(http.StatusNotFound, fmt.Sprintf("%d not found", id))
	}

	return c.JSON(http.StatusOK, todoList.Todos[id-1])
}

func main() {
	// we setup a global variable here to keep track of our data in memory
	// in the real world you would be pulling this data from a database or other microservice
	todoList = &TodoList{
		Title: "TODO List",
		Todos: []Todo{
			{ID: 1, Item: "Install GO", IsDone: true},
			{ID: 2, Item: "Create Microservice", IsDone: false},
		},
	}

	// We will use echo as our web framework
	// you can learn more at: https://echo.labstack.com/
	e := echo.New()
	// we don't need to have the echo banner show up every time
	e.HideBanner = true
	// set the log level to debug so we can see all messages
	e.Logger.SetLevel(log.DEBUG)
	// change the logger to be human readable
	e.Logger.SetHeader("${time_rfc3339} ${level}\t")

	e.Logger.Info("Getting templates")
	indexTemplate = template.Must(template.ParseFiles("templates/index.gohtml"))

	// e.Logger.Info("some info")
	// e.Logger.Debug("Some debug")
	// e.Logger.Error("some error")
	// e.Logger.Warn("some warn")

	// share our static resources / files
	e.Static("/resources", "resources")

	e.GET("/", getIndex)
	e.GET("/status", getStatus)
	e.GET("/todos", getTodos)
	e.GET("/todo/:id", getTodo)

	// get the port definition from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		//log.Print("No port specified, defaulting to " + port)
	} else {
		//log.Print("Using port " + port + " from environment")
	}

	// start the server
	e.Logger.Info("Starting server on port " + port + " ...")
	e.Logger.Fatal(e.Start(":" + port))
}
