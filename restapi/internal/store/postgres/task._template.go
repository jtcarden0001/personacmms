package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var taskTemplateTableName = "tasktemplate"

func (pg *Store) CreateTaskTemplate(taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	taskTemplate.Id = uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description, type) VALUES ($1, $2, $3, $4)`, taskTemplateTableName)
	_, err := pg.db.Exec(query, taskTemplate.Id, taskTemplate.Title, taskTemplate.Description, taskTemplate.Type)
	if err != nil {
		return tp.TaskTemplate{}, handleDbError(err, "task-template")
	}

	return taskTemplate, nil
}

func (pg *Store) DeleteTaskTemplate(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, taskTemplateTableName)
	_, err := pg.db.Exec(query, title)

	return handleDbError(err, "task-template")
}

func (pg *Store) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s`, taskTemplateTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "task-template")
	}
	defer rows.Close()

	var taskTemplates = []tp.TaskTemplate{}
	for rows.Next() {
		var taskTemplate tp.TaskTemplate
		err = rows.Scan(&taskTemplate.Id, &taskTemplate.Title, &taskTemplate.Description, &taskTemplate.Type)
		if err != nil {
			return nil, handleDbError(err, "task-template")
		}
		taskTemplates = append(taskTemplates, taskTemplate)
	}

	return taskTemplates, nil
}

func (pg *Store) GetTaskTemplate(title string) (tp.TaskTemplate, error) {
	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s WHERE title = $1`, taskTemplateTableName)
	row := pg.db.QueryRow(query, title)

	var taskTemplate tp.TaskTemplate
	err := row.Scan(&taskTemplate.Id, &taskTemplate.Title, &taskTemplate.Description, &taskTemplate.Type)
	if err != nil {
		return tp.TaskTemplate{}, handleDbError(err, "task-template")
	}

	return taskTemplate, nil
}

func (pg *Store) UpdateTaskTemplate(title string, taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2, type = $3 WHERE title = $4 returning id`, taskTemplateTableName)
	err := pg.db.QueryRow(query, taskTemplate.Title, taskTemplate.Description, taskTemplate.Type, taskTemplate.Title).Scan(&taskTemplate.Id)
	if err != nil {
		return tp.TaskTemplate{}, handleDbError(err, "task-template")
	}

	return taskTemplate, nil
}
