package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type WorkOrderStatus interface {
	CreateWorkOrderStatus(string) (int, error)
	DeleteWorkOrderStatus(int) error
	GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(int) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(int, string) error
}

func (pg *Store) CreateWorkOrderStatus(title string) (int, error) {
	query := `INSERT INTO work_order_status (title) VALUES ($1) RETURNING id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteWorkOrderStatus(id int) error {
	query := `DELETE FROM work_order_status WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	query := `SELECT id, title FROM work_order_status`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []tp.WorkOrderStatus
	for rows.Next() {
		var status tp.WorkOrderStatus
		err = rows.Scan(&status.Id, &status.Title)
		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, err
}

func (pg *Store) GetWorkOrderStatus(id int) (tp.WorkOrderStatus, error) {
	query := `SELECT id, title FROM work_order_status WHERE id = $1`
	var status tp.WorkOrderStatus
	err := pg.db.QueryRow(query, id).Scan(&status.Id, &status.Title)

	return status, err
}

func (pg *Store) UpdateWorkOrderStatus(id int, name string) error {
	query := `UPDATE work_order_status SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, name, id)

	return err
}
