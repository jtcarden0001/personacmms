package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var workOrderTable = "workorder"

func (pg *PostgresStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	wosId, err := pg.getWorkOrderStatusIdFromTitle(wo.Status)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	query := fmt.Sprintf(`INSERT INTO %s (
		id, title, created_date, completed_date, instructions, notes, cumulative_miles, cumulative_hours, status_id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)`, workOrderTable)

	_, err = pg.db.Exec(
		query,
		wo.Id,
		wo.Title,
		wo.CreatedDate,
		wo.CompletedDate,
		wo.Instructions,
		wo.Notes,
		wo.CumulativeMiles,
		wo.CumulativeHours,
		wosId,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	return wo, nil
}

func (pg *PostgresStore) DeleteWorkOrder(woId uuid.UUID) error {
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

func (pg *PostgresStore) GetWorkOrder(woId uuid.UUID) (tp.WorkOrder, error) {
	var wo tp.WorkOrder
	var statusId uuid.UUID
	query := fmt.Sprintf(`SELECT (
	  id, title, created_date, completed_date, instructions, notes, cumulative_miles, cumulative_hours, status_id
	  ) FROM %s WHERE id = $1`, workOrderTable)
	err := pg.db.QueryRow(
		query,
		woId).Scan(
		&wo.Id,
		&wo.Title,
		&wo.CreatedDate,
		&wo.CompletedDate,
		&wo.Instructions,
		&wo.Notes,
		&wo.CumulativeMiles,
		&wo.CumulativeHours,
		&statusId,
	)
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	statusTitle, err := pg.getWorkOrderStatusTitleFromId(statusId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	wo.Status = statusTitle
	return wo, nil
}

func (pg *PostgresStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	query := fmt.Sprintf(`SELECT (
	  id, title, created_date, completed_date, instructions, notes, cumulative_miles, cumulative_hours, status_id
	  ) FROM %s`, workOrderTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "work-order")
	}

	workOrders := []tp.WorkOrder{}
	for rows.Next() {
		var wo tp.WorkOrder
		var statusId uuid.UUID
		err = rows.Scan(
			&wo.Id,
			&wo.Title,
			&wo.CreatedDate,
			&wo.CompletedDate,
			&wo.Instructions,
			&wo.Notes,
			&wo.CumulativeMiles,
			&wo.CumulativeHours,
			&statusId,
		)
		if err != nil {
			return nil, handleDbError(err, "work-order")
		}

		statusTitle, err := pg.getWorkOrderStatusTitleFromId(statusId)
		if err != nil {
			return nil, err
		}

		wo.Status = statusTitle
		workOrders = append(workOrders, wo)
	}

	return workOrders, nil
}

func (pg *PostgresStore) UpdateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	wosId, err := pg.getWorkOrderStatusIdFromTitle(wo.Status)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	query := fmt.Sprintf(`UPDATE %s SET 
		title = $1
		created_date = $2, 
		completed_date = $3,
		instructions = $4,
		notes = $5,
		cumulative_miles = $6,
		cumulative_hours = $7,
		status_id = $8
		WHERE id = $9`, workOrderTable)

	result, err := pg.db.Exec(
		query,
		wo.Title,
		wo.CreatedDate,
		wo.CompletedDate,
		wo.Instructions,
		wo.Notes,
		wo.CumulativeMiles,
		wo.CumulativeHours,
		wosId,
		wo.Id,
	)

	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.WorkOrder{}, handleDbError(err, "work-order")
	}

	if rowsAffected == 0 {
		return tp.WorkOrder{}, errors.Wrapf(ae.ErrNotFound, "work-order with id %s not found", wo.Id)
	}

	return wo, nil
}

var woStatusTableName = "workorderstatus"

func (pg *PostgresStore) getWorkOrderStatusTitleFromId(id uuid.UUID) (string, error) {
	query := fmt.Sprintf(`SELECT title FROM %s WHERE id = $1`, woStatusTableName)
	var title string
	err := pg.db.QueryRow(query, id).Scan(&title)
	if err != nil {
		return "", handleDbError(err, "status")
	}

	return title, nil
}

func (pg *PostgresStore) getWorkOrderStatusIdFromTitle(title string) (uuid.UUID, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE title = $1`, woStatusTableName)
	var id uuid.UUID
	err := pg.db.QueryRow(query, title).Scan(&id)
	if err != nil {
		return uuid.Nil, handleDbError(err, "status")
	}

	return id, nil
}
