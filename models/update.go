package models

import "todo-list-go/db"

func UpdateTodo(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(
		`UPDATE TODOS 
		SET title = $1,
		description = $2,
		done = $3
		WHERE id = $4
		`,
		todo.Title,
		todo.Description,
		todo.Done,
		id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
