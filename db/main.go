package db

import (
	"database/sql"
	"fmt"

	"github.com/KrishGarg/go-todo-api/model"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	SqlDB *sql.DB
}

var database Database

func Connect() error {
	conn, err := sql.Open("sqlite3", "./file.db")

	if err != nil {
		return err
	}

	database.SqlDB = conn
	fmt.Println("Connected with the database!")
	return nil
}

func GetDB() *Database {
	return &database
}

func (db *Database) Prepare() error {
	_, err := db.SqlDB.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		todo TEXT NOT NULL,
		done INTEGER NOT NULL DEFAULT 0
	)`)

	if err != nil {
		return err
	}

	fmt.Println("Prepared the database!")
	return nil
}

func (db *Database) GetTodos() ([]model.Todo, error) {
	rows, err := db.SqlDB.Query("SELECT id, todo, done FROM todos")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo

	for rows.Next() {
		var (
			id   int
			todo string
			done int
		)
		if err := rows.Scan(&id, &todo, &done); err != nil {
			return nil, err
		}

		var todoBool bool

		if done == 1 {
			todoBool = true
		} else {
			todoBool = false
		}

		todoStruct := model.Todo{
			Id:   id,
			Todo: todo,
			Done: todoBool,
		}

		todos = append(todos, todoStruct)
	}

	return todos, nil
}

func (db *Database) AddTodo(todo string) (*model.Todo, error) {
	result, err := db.SqlDB.Exec("INSERT INTO todos (todo) VALUES (?)", todo)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	todoStruct := model.Todo{
		Id:   int(id),
		Todo: todo,
		Done: false,
	}
	return &todoStruct, nil
}

func (db *Database) ToggleTodo(id int) (*model.Todo, error) {

	row := db.SqlDB.QueryRow("SELECT todo, done FROM todos WHERE id = ?", id)

	var (
		todo string
		done int
	)
	if err := row.Scan(&todo, &done); err != nil {
		return nil, err
	}

	var doneToggledInt int

	if done == 1 {
		doneToggledInt = 0
	} else {
		doneToggledInt = 1
	}

	if _, err := db.SqlDB.Exec("UPDATE todos SET done = ? WHERE id = ?", doneToggledInt, id); err != nil {
		return nil, err
	}

	var doneBool bool

	if doneToggledInt == 1 {
		doneBool = true
	} else {
		doneBool = false
	}

	todoStruct := model.Todo{
		Id:   int(id),
		Todo: todo,
		Done: doneBool,
	}
	return &todoStruct, nil
}

func (db *Database) GetTodoByID(id int) (*model.Todo, error) {
	row := db.SqlDB.QueryRow("SELECT todo, done FROM todos WHERE id = ?", id)

	var (
		todo string
		done int
	)
	if err := row.Scan(&todo, &done); err != nil {
		return nil, err
	}

	var doneBool bool

	if done == 1 {
		doneBool = true
	} else {
		doneBool = false
	}

	todoStruct := model.Todo{
		Id:   int(id),
		Todo: todo,
		Done: doneBool,
	}
	return &todoStruct, nil
}
