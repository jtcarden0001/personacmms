package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Task interface {
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(string) error
	ListTasks() ([]tp.Task, error)
	GetTask(string) (tp.Task, error)
	UpdateTask(string, tp.Task) (tp.Task, error)
}

var taskTableName = "task"

func (pg *Store) CreateTask(task tp.Task) (tp.Task, error) {
	task.Id = uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, taskTableName)
	_, err := pg.db.Exec(query, task.Id.String(), task.Title, task.Description)
	if err != nil {
		return tp.Task{}, err
	}

	return task, nil
}

func (pg *Store) DeleteTask(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, taskTableName)
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListTasks() ([]tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s`, taskTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks = []tp.Task{}
	for rows.Next() {
		var task tp.Task
		err = rows.Scan(&task.Id, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (pg *Store) GetTask(title string) (tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, taskTableName)
	row := pg.db.QueryRow(query, title)

	var task tp.Task
	err := row.Scan(&task.Id, &task.Title, &task.Description)
	if err != nil {
		return tp.Task{}, err
	}

	return task, nil
}

func (pg *Store) UpdateTask(title string, task tp.Task) (tp.Task, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3 returning id`, taskTableName)
	err := pg.db.QueryRow(query, task.Title, task.Description, title).Scan(&task.Id)
	if err != nil {
		return tp.Task{}, err
	}

	return task, nil
}
