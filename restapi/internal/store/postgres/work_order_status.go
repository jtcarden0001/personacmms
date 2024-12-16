package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/pkg/errors"
)

var workOrderStatusTableName = "workorderstatus"

func (s *Store) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2)`, workOrderStatusTableName)
	_, err := s.db.Exec(query, id, wos.Title)
	if err != nil {
		return tp.WorkOrderStatus{}, handleDbError(err, "work-order-status")
	}

	wos.Id = id
	return wos, nil
}

func (s *Store) DeleteWorkOrderStatus(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, workOrderStatusTableName)
	result, err := s.db.Exec(query, title)
	if err != nil {
		return handleDbError(err, "work-order-status")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "work-order-status")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "work-order-status with title %s not found", title)
	}
	return nil
}

func (s *Store) ListWorkOrderStatuses() ([]tp.WorkOrderStatus, error) {
	var workOrderStatuses = []tp.WorkOrderStatus{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, workOrderStatusTableName)
	rows, err := s.db.Query(query)
	if err != nil {
		return workOrderStatuses, handleDbError(err, "work-order-status")
	}
	defer rows.Close()

	for rows.Next() {
		var wos tp.WorkOrderStatus
		err = rows.Scan(&wos.Id, &wos.Title)
		if err != nil {
			return nil, handleDbError(err, "work-order-status")
		}
		workOrderStatuses = append(workOrderStatuses, wos)
	}

	return workOrderStatuses, nil
}

func (s *Store) GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error) {
	var wos tp.WorkOrderStatus
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE title = $1`, workOrderStatusTableName)
	err := s.db.QueryRow(query, title).Scan(&wos.Id, &wos.Title)
	if err != nil {
		return tp.WorkOrderStatus{}, handleDbError(err, "work-order-status")
	}

	return wos, nil
}

func (s *Store) UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2 RETURNING id`, workOrderStatusTableName)
	err := s.db.QueryRow(query, wos.Title, title).Scan(&wos.Id)
	if err != nil {
		return tp.WorkOrderStatus{}, handleDbError(err, "work-order-status")
	}

	return wos, nil
}
