package postgres

import (
	"database/sql"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type TaskConsumable interface {
	CreateTaskConsumable(int, int, string) error
	DeleteTaskConsumable(int, int) error
	GetAllTaskConsumable() ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByTaskId(int) ([]tp.TaskConsumable, error)
	GetTaskConsumable(int, int) (tp.TaskConsumable, error)
}

func (pg *Store) CreateTaskConsumable(taskId int, consumableId int, quantity string) error {
	query := `INSERT INTO task_consumable (task_id, consumable_id, quantity_note) VALUES ($1, $2, $3)`
	_, err := pg.db.Exec(query, taskId, consumableId, quantity)

	return err
}

func (pg *Store) DeleteTaskConsumable(taskId int, consumableId int) error {
	query := `DELETE FROM task_consumable WHERE task_id = $1 AND consumable_id = $2`
	_, err := pg.db.Exec(query, taskId, consumableId)

	return err
}

func (pg *Store) GetAllTaskConsumable() ([]tp.TaskConsumable, error) {
	query := `SELECT task_id, consumable_id, quantity_note FROM task_consumable`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskConsumableList(rows)
}

func (pg *Store) GetAllTaskConsumableByTaskId(taskId int) ([]tp.TaskConsumable, error) {
	query := `SELECT task_id, consumable_id, quantity_note FROM task_consumable WHERE task_id = $1`
	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskConsumableList(rows)
}

func (pg *Store) GetTaskConsumable(consumableId int, taskId int) (tp.TaskConsumable, error) {
	query := `SELECT task_id, consumable_id, quantity_note FROM task_consumable WHERE task_id = $1 AND consumable_id = $2`
	var tc tp.TaskConsumable
	err := pg.db.QueryRow(query, taskId, consumableId).Scan(&tc.TaskId, &tc.ConsumableId, &tc.QuantityNote)

	return tc, err
}

func populateTaskConsumableList(rows *sql.Rows) ([]tp.TaskConsumable, error) {
	var taskConsumables []tp.TaskConsumable
	for rows.Next() {
		var tc tp.TaskConsumable
		err := rows.Scan(&tc.TaskId, &tc.ConsumableId, &tc.QuantityNote)
		if err != nil {
			return nil, err
		}
		taskConsumables = append(taskConsumables, tc)
	}

	return taskConsumables, nil
}
