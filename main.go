package main

import (
	"log"

	"github.com/KrishGarg/go-todo-api/db"
	"github.com/KrishGarg/go-todo-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := db.Connect(); err != nil {
		panic(err)
	}

	database := db.GetDB()
	defer database.SqlDB.Close()

	database.Prepare()

	app := fiber.New()

	api := app.Group("/api")
	todos := api.Group("/todos")

	todos.Get("/", handlers.GetTodos)
	todos.Post("/", handlers.AddTodo)
	todos.Patch("/", handlers.ToggleTodo)

	log.Fatal(app.Listen(":3000"))
}
