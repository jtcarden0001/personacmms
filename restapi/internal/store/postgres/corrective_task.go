package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type CorrectiveTask interface {
	CreateCorrectiveTask(tp.CorrectiveTask) (tp.CorrectiveTask, error)
	DeleteCorrectiveTask(string) error
	ListCorrectiveTasks() ([]tp.CorrectiveTask, error)
	GetCorrectiveTask(string) (tp.CorrectiveTask, error)
	UpdateCorrectiveTask(string, tp.CorrectiveTask) (tp.CorrectiveTask, error)
}

var correctiveTaskTableName = "correctivetask"

func (pg *Store) CreateCorrectiveTask(correctiveTask tp.CorrectiveTask) (tp.CorrectiveTask, error) {
	correctiveTask.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, correctiveTaskTableName)
	_, err := pg.db.Exec(query, correctiveTask.Id, correctiveTask.Title, correctiveTask.Description)
	if err != nil {
		return tp.CorrectiveTask{}, err
	}

	return correctiveTask, nil
}

func (pg *Store) DeleteCorrectiveTask(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, correctiveTaskTableName)
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListCorrectiveTasks() ([]tp.CorrectiveTask, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s`, correctiveTaskTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var correctiveTasks = []tp.CorrectiveTask{}
	for rows.Next() {
		var correctiveTask tp.CorrectiveTask
		err = rows.Scan(&correctiveTask.Id, &correctiveTask.Title, &correctiveTask.Description)
		if err != nil {
			return nil, err
		}
		correctiveTasks = append(correctiveTasks, correctiveTask)
	}

	return correctiveTasks, nil
}

func (pg *Store) GetCorrectiveTask(title string) (tp.CorrectiveTask, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, correctiveTaskTableName)
	row := pg.db.QueryRow(query, title)

	var correctiveTask tp.CorrectiveTask
	err := row.Scan(&correctiveTask.Id, &correctiveTask.Title, &correctiveTask.Description)
	if err != nil {
		return tp.CorrectiveTask{}, err
	}

	return correctiveTask, nil
}

func (pg *Store) UpdateCorrectiveTask(oldTitle string, correctiveTask tp.CorrectiveTask) (tp.CorrectiveTask, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3`, correctiveTaskTableName)
	_, err := pg.db.Exec(query, correctiveTask.Title, correctiveTask.Description, oldTitle)
	if err != nil {
		return tp.CorrectiveTask{}, err
	}

	return correctiveTask, nil
}
