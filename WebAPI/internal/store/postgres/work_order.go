package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(int, int, string, *string) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByEquipmentId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, int, string, *string) error
}

func (pg *Store) CreateWorkOrder(taskId int, statusId int, createdDateTime string, CompleteDateTime *string) (int, error) {
	query := `INSERT INTO work_order (task_id, status_id, create_date, complete_date) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := pg.db.QueryRow(query, taskId, statusId, createdDateTime, CompleteDateTime).Scan(&id)

	return id, err
}

func (pg *Store) DeleteWorkOrder(id int) error {
	query := `DELETE FROM work_order WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllWorkOrder() ([]tp.WorkOrder, error) {
	query := `SELECT id, task_id, status_id, create_date, complete_date FROM work_order`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []tp.WorkOrder
	for rows.Next() {
		var wo tp.WorkOrder
		err = rows.Scan(&wo.Id, &wo.TaskId, &wo.StatusId, &wo.CreatedDate, &wo.CompletedDate)
		if err != nil {
			return nil, err
		}

		orders = append(orders, wo)
	}

	return orders, err
}

func (pg *Store) GetAllWorkOrderByEquipmentId(equipmentId int) ([]tp.WorkOrder, error) {
	return []tp.WorkOrder{}, errors.New("not implemented")
}

func (pg *Store) GetWorkOrder(id int) (tp.WorkOrder, error) {
	query := `SELECT id, task_id, status_id, create_date, complete_date FROM work_order WHERE id = $1`
	var wo tp.WorkOrder
	err := pg.db.QueryRow(query, id).Scan(&wo.Id, &wo.TaskId, &wo.StatusId, &wo.CreatedDate, &wo.CompletedDate)

	return wo, err
}

func (pg *Store) UpdateWorkOrder(id int, taskId int, statusId int, startDateTime string, CompleteDateTime *string) error {
	query := `UPDATE work_order SET task_id = $1, status_id = $2, create_date = $3, complete_date = $4 WHERE id = $5`
	_, err := pg.db.Exec(query, taskId, statusId, startDateTime, CompleteDateTime, id)

	return err
}
