package models

import "todo-list-go/db"

func DeleteTodo(id int64) (int64, error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(
		`DELETE FROM TODOS
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
