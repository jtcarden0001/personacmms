package postgres

import (
	"database/sql"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Task interface {
	CreateTask(string, string, *int, *int, *int, *int, int) (int, error)
	DeleteTask(int) error
	GetAllTask() ([]tp.Task, error)
	GetAllTaskByEquipmentId(int) ([]tp.Task, error)
	GetTask(int) (tp.Task, error)
	UpdateTask(int, string, string, *int, *int, *int, *int, int) error
}

func (pg *Store) CreateTask(title string, instructions string, timeQuant *int, timeUnitId *int, usageQuant *int, usageUnitId *int, equipmentId int) (int, error) {
	query := `INSERT INTO task (title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`
	var id int
	err := pg.db.QueryRow(query, title, instructions, timeQuant, timeUnitId, usageQuant, usageUnitId, equipmentId).Scan(&id)

	return id, err
}

func (pg *Store) DeleteTask(id int) error {
	query := `DELETE FROM task WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllTask() ([]tp.Task, error) {
	query := `SELECT id, title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id FROM task`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskList(rows)
}

func (pg *Store) GetAllTaskByEquipmentId(equipmentId int) ([]tp.Task, error) {
	query := `SELECT id, title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id FROM task WHERE equipment_id = $1`
	rows, err := pg.db.Query(query, equipmentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskList(rows)
}

func (pg *Store) GetTask(id int) (tp.Task, error) {
	query := `SELECT id, title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id FROM task WHERE id = $1`
	var t tp.Task
	err := pg.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Instructions, &t.TimePeriodicityQuantity, &t.TimePeriodicityUnitId, &t.UsagePeriodicityQuantity, &t.UsagePeriodicityUnitId, &t.EquipmentId)

	return t, err
}

func (pg *Store) UpdateTask(id int, title string, instructions string, timeQuant *int, timeUnit *int, usageQuant *int, usageUnit *int, equipmentId int) error {
	query := `UPDATE task SET title = $1, instructions = $2, time_periodicity_quantity = $3, time_periodicity_unit_id = $4, usage_periodicity_quantity = $5, usage_periodicity_unit_id = $6, equipment_id = $7 WHERE id = $8`
	_, err := pg.db.Exec(query, title, instructions, timeQuant, timeUnit, usageQuant, usageUnit, equipmentId, id)

	return err
}

func populateTaskList(rows *sql.Rows) ([]tp.Task, error) {
	var tasks []tp.Task
	for rows.Next() {
		var t tp.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Instructions, &t.TimePeriodicityQuantity, &t.TimePeriodicityUnitId, &t.UsagePeriodicityQuantity, &t.UsagePeriodicityUnitId, &t.EquipmentId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
