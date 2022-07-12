package models

import "todo-list-go/db"

func InsertTodo(todo Todo) (id int64, err error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	err = conn.QueryRow(
		`insert into todos (title, description, done)
		values ($1, $2, $3)
		returning id
		`,
		todo.Title,
		todo.Description,
		todo.Done,
	).Scan(&id)

	return
}
