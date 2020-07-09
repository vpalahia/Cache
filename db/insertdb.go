package db

import (
	"fmt"

	"github.com/vpalahia/Cache/types"
)

func InsertIntoDB(todoitem types.Todo) error {
	insertQuery := `INSERT INTO todotasks (id, task, description) VALUES ($1, $2, $3)`
	_, err := db.Exec(insertQuery, todoitem.Id, todoitem.Task, todoitem.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
