package postgres

import (
	"errors"

	tm "time"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(int, int, tm.Time, *tm.Time) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByAssetId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, tm.Time, *tm.Time) error
}

func (pg *Store) CreateWorkOrder(preventativeTaskId int, statusId int, createdDateTime tm.Time, CompleteDateTime *tm.Time) (int, error) {
	query := `INSERT INTO work_order (preventativeTask_id, status_id, create_date, complete_date) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := pg.db.QueryRow(query, preventativeTaskId, statusId, createdDateTime, CompleteDateTime).Scan(&id)

	return id, err
}

func (pg *Store) DeleteWorkOrder(id int) error {
	query := `DELETE FROM work_order WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllWorkOrder() ([]tp.WorkOrder, error) {
	query := `SELECT id, preventativeTask_id, status_id, create_date, complete_date FROM work_order`
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

func (pg *Store) GetAllWorkOrderByAssetId(assetId int) ([]tp.WorkOrder, error) {
	return []tp.WorkOrder{}, errors.New("not implemented")
}

func (pg *Store) GetWorkOrder(id int) (tp.WorkOrder, error) {
	query := `SELECT id, preventativeTask_id, status_id, create_date, complete_date FROM work_order WHERE id = $1`
	var wo tp.WorkOrder
	err := pg.db.QueryRow(query, id).Scan(&wo.Id, &wo.TaskId, &wo.StatusId, &wo.CreatedDate, &wo.CompletedDate)

	return wo, err
}

func (pg *Store) UpdateWorkOrder(id int, statusId int, startDateTime tm.Time, CompleteDateTime *tm.Time) error {
	// A work order is only associated with 1 preventativeTask, the ability to update the preventativeTaskId on a work order doesn't make sense
	// as it is just a status descriptor to a preventativeTask that needs to be completed
	query := `UPDATE work_order SET status_id = $1, create_date = $2, complete_date = $3 WHERE id = $4`
	_, err := pg.db.Exec(query, statusId, startDateTime, CompleteDateTime, id)

	return err
}
