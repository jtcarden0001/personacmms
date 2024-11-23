package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type WorkOrderStatus interface {
	CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(title string) error
	ListWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
}

var workOrderStatusTableName = "workorderstatus"

func (s *Store) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2)`, workOrderStatusTableName)
	_, err := s.db.Exec(query, id, wos.Title)
	if err != nil {
		return tp.WorkOrderStatus{}, processDbError(err)
	}

	wos.Id = id
	return wos, nil
}

func (s *Store) DeleteWorkOrderStatus(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, workOrderStatusTableName)
	_, err := s.db.Exec(query, title)

	return processDbError(err)
}

func (s *Store) ListWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	var workOrderStatuses = []tp.WorkOrderStatus{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, workOrderStatusTableName)
	rows, err := s.db.Query(query)
	if err != nil {
		return workOrderStatuses, processDbError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var wos tp.WorkOrderStatus
		err = rows.Scan(&wos.Id, &wos.Title)
		if err != nil {
			return nil, err
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
		return tp.WorkOrderStatus{}, processDbError(err)
	}

	return wos, nil
}

func (s *Store) UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2`, workOrderStatusTableName)
	_, err := s.db.Exec(query, wos.Title, title)
	if err != nil {
		return tp.WorkOrderStatus{}, processDbError(err)
	}

	return wos, nil
}
