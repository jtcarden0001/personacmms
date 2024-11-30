package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var taskTableName = "task"

func (pg *Store) CreateTask(task tp.Task) (tp.Task, error) {
	task.Id = uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description, type) VALUES ($1, $2, $3, $4)`, taskTableName)
	_, err := pg.db.Exec(query, task.Id, task.Title, task.Description, task.Type)
	if err != nil {
		return tp.Task{}, handleDbError(err, "task")
	}

	return task, nil
}

func (pg *Store) DeleteTask(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, taskTableName)
	_, err := pg.db.Exec(query, title)

	return handleDbError(err, "task")
}

func (pg *Store) ListTasks() ([]tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s`, taskTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "task")
	}
	defer rows.Close()

	var tasks = []tp.Task{}
	for rows.Next() {
		var task tp.Task
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Type)
		if err != nil {
			return nil, handleDbError(err, "task")
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (pg *Store) GetTask(title string) (tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s WHERE title = $1`, taskTableName)
	row := pg.db.QueryRow(query, title)

	var task tp.Task
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Type)
	if err != nil {
		return tp.Task{}, handleDbError(err, "task")
	}

	return task, nil
}

func (pg *Store) UpdateTask(title string, task tp.Task) (tp.Task, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2, type = $3 WHERE title = $4 returning id`, taskTableName)
	err := pg.db.QueryRow(query, task.Title, task.Description, task.Type, task.Title).Scan(&task.Id)
	if err != nil {
		return tp.Task{}, handleDbError(err, "task")
	}

	return task, nil
}
