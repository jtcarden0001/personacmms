package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

type WorkOrder interface {
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(tp.UUID) error
	ListTaskWorkOrders(tp.UUID) ([]tp.WorkOrder, error)
	GetWorkOrder(tp.UUID) (tp.WorkOrder, error)
	UpdateWorkOrder(tp.UUID, tp.WorkOrder) (tp.WorkOrder, error)
}

var workOrderTable = "workorder"

func (pg *PostgresStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	wo.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (
		id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, task_id, status_title
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
		wo.TaskId,
		wo.StatusTitle,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

func (pg *PostgresStore) DeleteWorkOrder(woId tp.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, workOrderTable)
	result, err := pg.db.Exec(query, woId)
	if err != nil {
		return handleDbError(err, "work-order")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "work-order")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "work-order with id %s not found", woId)
	}

	return nil
}

func (pg *PostgresStore) GetWorkOrder(woId tp.UUID) (tp.WorkOrder, error) {
	var wo tp.WorkOrder
	query := fmt.Sprintf(`SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, task_id, status_title FROM %s WHERE id = $1`, workOrderTable)
	err := pg.db.QueryRow(query, woId).Scan(&wo.Id, &wo.CreatedDate, &wo.CompletedDate, &wo.Notes, &wo.CumulativeMiles, &wo.CumulativeHours, &wo.TaskId, &wo.StatusTitle)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

// TODO: add testing for this
func (pg *PostgresStore) GetWorkOrderForTask(tId tp.UUID, woId tp.UUID) (tp.WorkOrder, error) {
	var wo tp.WorkOrder
	query := fmt.Sprintf(`
		SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, task_id, status_title
		FROM %s
		WHERE id = $1 AND task_id = $2
	`, workOrderTable)

	row := pg.db.QueryRow(query, woId, tId)
	err := row.Scan(
		&wo.Id,
		&wo.CreatedDate,
		&wo.CompletedDate,
		&wo.Notes,
		&wo.CumulativeMiles,
		&wo.CumulativeHours,
		&wo.TaskId,
		&wo.StatusTitle,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

func (pg *PostgresStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	query := fmt.Sprintf(`SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, task_id, status_title FROM %s`, workOrderTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "work-order")
	}

	workOrders := []tp.WorkOrder{}
	for rows.Next() {
		var wo tp.WorkOrder
		err = rows.Scan(&wo.Id, &wo.CreatedDate, &wo.CompletedDate, &wo.Notes, &wo.CumulativeMiles, &wo.CumulativeHours, &wo.TaskId, &wo.StatusTitle)
		if err != nil {
			return nil, handleDbError(err, "work-order")
		}
		workOrders = append(workOrders, wo)
	}

	return workOrders, nil
}

// TODO: add testing for this
func (pg *PostgresStore) ListWorkOrdersByTaskId(tId tp.UUID) ([]tp.WorkOrder, error) {
	query := fmt.Sprintf(`
		SELECT id, created_date, completed_date, notes, cumulative_miles, cumulative_hours, task_id, status_title
		FROM %s
		WHERE task_id = $1
	`, workOrderTable)

	rows, err := pg.db.Query(query, tId)
	if err != nil {
		return nil, handleDbError(err, "work-order")
	}

	workOrders := []tp.WorkOrder{}
	for rows.Next() {
		var wo tp.WorkOrder
		err = rows.Scan(&wo.Id, &wo.CreatedDate, &wo.CompletedDate,
			&wo.Notes, &wo.CumulativeMiles, &wo.CumulativeHours, &wo.TaskId, &wo.StatusTitle)
		if err != nil {
			return nil, handleDbError(err, "work-order")
		}

		workOrders = append(workOrders, wo)
	}

	return workOrders, nil
}

func (pg *PostgresStore) UpdateWorkOrder(woId tp.UUID, wo tp.WorkOrder) (tp.WorkOrder, error) {
	query := fmt.Sprintf(`UPDATE %s SET 
		created_date = $2, 
		completed_date = $3, 
		notes = $4, 
		cumulative_miles = $5, 
		cumulative_hours = $6, 
		task_id = $7, 
		status_title = $8 
		WHERE id = $1`, workOrderTable)

	result, err := pg.db.Exec(
		query,
		woId,
		wo.CreatedDate,
		wo.CompletedDate,
		wo.Notes,
		wo.CumulativeMiles,
		wo.CumulativeHours,
		wo.TaskId,
		wo.StatusTitle,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	if rowsAffected == 0 {
		return tp.WorkOrder{}, errors.Wrapf(ae.ErrNotFound, "work-order with id %s not found", woId)
	}

	return wo, nil
}
