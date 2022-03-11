package handlers

import (
	"strconv"

	"github.com/KrishGarg/go-todo-api/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var database *db.Database = db.GetDB()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(stru interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(stru)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func GetTodos(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		todos, err := database.GetTodos()
		if err != nil {
			return err
		}
		return c.JSON(todos)
	} else {
		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			return err
		}
		todo, err := database.GetTodoByID(int(idInt))
		return c.JSON(todo)
	}
}

func AddTodo(c *fiber.Ctx) error {

	type AddTodoRequestBody struct {
		Todo string `json:"todo" validate:"required,min=5"`
	}

	todo := new(AddTodoRequestBody)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	errors := ValidateStruct(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	todoStruct, err := database.AddTodo(todo.Todo)
	if err != nil {
		return err
	}

	return c.JSON(todoStruct)
}

func ToggleTodo(c *fiber.Ctx) error {
	type ToggleTodoRequestBody struct {
		Id int `json:"id" validate:"required,number"`
	}

	todo := new(ToggleTodoRequestBody)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	errors := ValidateStruct(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	todoStruct, err := database.ToggleTodo(todo.Id)
	if err != nil {
		return err
	}

	return c.JSON(todoStruct)
}
