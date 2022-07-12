package models

import "todo-list-go/db"

func GetAllTodos() (todos []Todo, err error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	query := "SELECT * FROM TODOS"

	rows, err := conn.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		todo := Todo{}

		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done); err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}

func GetById(id int64) (todo Todo, err error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	queryString := "SELECT * FROM TODOS WHERE id = $1"

	row := conn.QueryRow(queryString, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
