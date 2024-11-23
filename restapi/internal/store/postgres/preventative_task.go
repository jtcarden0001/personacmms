package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type PreventativeTask interface {
	CreatePreventativeTask(tp.PreventativeTask) (tp.PreventativeTask, error)
	DeletePreventativeTask(string) error
	ListPreventativeTasks() ([]tp.PreventativeTask, error)
	GetPreventativeTask(string) (tp.PreventativeTask, error)
	UpdatePreventativeTask(string, tp.PreventativeTask) (tp.PreventativeTask, error)
}

var preventativeTaskTableName = "preventativeTask"

func (pg *Store) CreatePreventativeTask(preventativeTask tp.PreventativeTask) (tp.PreventativeTask, error) {
	preventativeTask.Id = uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, preventativeTaskTableName)
	_, err := pg.db.Exec(query, preventativeTask.Id, preventativeTask.Title, preventativeTask.Description)
	if err != nil {
		return tp.PreventativeTask{}, err
	}

	return preventativeTask, nil
}

func (pg *Store) DeletePreventativeTask(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, preventativeTaskTableName)
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListPreventativeTasks() ([]tp.PreventativeTask, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s`, preventativeTaskTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var preventativeTasks = []tp.PreventativeTask{}
	for rows.Next() {
		var preventativeTask tp.PreventativeTask
		err = rows.Scan(&preventativeTask.Id, &preventativeTask.Title, &preventativeTask.Description)
		if err != nil {
			return nil, err
		}
		preventativeTasks = append(preventativeTasks, preventativeTask)
	}

	return preventativeTasks, nil
}

func (pg *Store) GetPreventativeTask(title string) (tp.PreventativeTask, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, preventativeTaskTableName)
	row := pg.db.QueryRow(query, title)

	var preventativeTask tp.PreventativeTask
	err := row.Scan(&preventativeTask.Id, &preventativeTask.Title, &preventativeTask.Description)
	if err != nil {
		return tp.PreventativeTask{}, err
	}

	return preventativeTask, nil
}

func (pg *Store) UpdatePreventativeTask(title string, preventativeTask tp.PreventativeTask) (tp.PreventativeTask, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3 returning id`, preventativeTaskTableName)
	err := pg.db.QueryRow(query, preventativeTask.Title, preventativeTask.Description, title).Scan(&preventativeTask.Id)
	if err != nil {
		return tp.PreventativeTask{}, err
	}

	return preventativeTask, nil
}
