package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(tp.UUID) error
	ListAssetTaskWorkOrders(tp.UUID) ([]tp.WorkOrder, error)
	GetWorkOrder(tp.UUID) (tp.WorkOrder, error)
	UpdateWorkOrder(tp.UUID, tp.WorkOrder) (tp.WorkOrder, error)
}

var workOrderTable = "workorder"

func (pg *Store) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	wo.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (
		id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, assettask_id, status_title
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	)`, workOrderTable)

	_, err := pg.db.Exec(
		query,
		wo.Id,
		wo.CreatedDate,
		wo.CompletedDate,
		wo.Notes,
		wo.CumulativeMiles,
		wo.CumulativeHours,
		wo.AssetTaskId,
		wo.StatusTitle,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

func (pg *Store) DeleteWorkOrder(woId tp.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, workOrderTable)
	_, err := pg.db.Exec(query, woId)
	return handleDbError(err, "work-order")
}

func (pg *Store) ListWorkOrders() ([]tp.WorkOrder, error) {
	query := fmt.Sprintf(`SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, assettask_id, status_title FROM %s`, workOrderTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "work-order")
	}

	workOrders := []tp.WorkOrder{}
	for rows.Next() {
		var wo tp.WorkOrder
		err = rows.Scan(&wo.Id, &wo.CreatedDate, &wo.CompletedDate, &wo.Notes, &wo.CumulativeMiles, &wo.CumulativeHours, &wo.AssetTaskId, &wo.StatusTitle)
		if err != nil {
			return nil, handleDbError(err, "work-order")
		}
		workOrders = append(workOrders, wo)
	}

	return workOrders, nil
}

func (pg *Store) GetWorkOrder(woId tp.UUID) (tp.WorkOrder, error) {
	var wo tp.WorkOrder
	query := fmt.Sprintf(`SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, assettask_id, status_title FROM %s WHERE id = $1`, workOrderTable)
	err := pg.db.QueryRow(query, woId).Scan(&wo.Id, &wo.CreatedDate, &wo.CompletedDate, &wo.Notes, &wo.CumulativeMiles, &wo.CumulativeHours, &wo.AssetTaskId, &wo.StatusTitle)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

func (pg *Store) UpdateWorkOrder(woId tp.UUID, wo tp.WorkOrder) (tp.WorkOrder, error) {
	query := fmt.Sprintf(`UPDATE %s SET 
		created_date = $2, 
		completed_date = $3, 
		notes = $4, 
		cumulative_miles = $5, 
		cumulative_hours = $6, 
		assettask_id = $7, 
		status_title = $8 
		WHERE id = $1`, workOrderTable)

	_, err := pg.db.Exec(
		query,
		woId,
		wo.CreatedDate,
		wo.CompletedDate,
		wo.Notes,
		wo.CumulativeMiles,
		wo.CumulativeHours,
		wo.AssetTaskId,
		wo.StatusTitle,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}
