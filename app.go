package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/friendsofgo/graphiql"
	"github.com/salvo1404/go-echo-graphql/db"
	echoGraphql "github.com/salvo1404/go-echo-graphql/graphql"
)

func main() {
	app := echo.New()

	// Config Logger
	app.Use(middleware.Logger())
	// 	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 		Format: "method=${method}, uri=${uri}, status=${status}\n",
	// 	}))

	fmt.Println("ciao")
	// Database
	d := db.Connect()
	d.LogMode(true)
	defer d.Close()

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go echo API !")
	})

	// Graphql endpoint here
	graphHandler, err := echoGraphql.NewGraphHandler(d)
	if err != nil {
		log.Fatalln(err)
	}
	app.POST("/graphql", echo.WrapHandler(graphHandler))

	// Playground Endpoint http://localhost:1444/graphiql
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Fatalln(err)
	}
	app.GET("/graphiql", echo.WrapHandler(graphiqlHandler))

	// Application start
	app.Logger.Fatal(app.Start(":1444"))
}
