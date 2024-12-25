package postgres

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
// 	"github.com/pkg/errors"
// )

// // moved out of mvp scope.  Will implement in future.

// var taskTemplateTableName = "tasktemplate"

// func (pg *PostgresStore) CreateTaskTemplate(tt tp.TaskTemplate) (tp.TaskTemplate, error) {
// 	taskTemplate.Id = uuid.New()
// 	query := fmt.Sprintf(`INSERT INTO %s (id, title, description, type) VALUES ($1, $2, $3, $4)`, taskTemplateTableName)
// 	_, err := pg.db.Exec(query, taskTemplate.Id, taskTemplate.Title, taskTemplate.Description, taskTemplate.Type)
// 	if err != nil {
// 		return tp.TaskTemplate{}, handleDbError(err, "task-template")
// 	}

// 	return taskTemplate, nil
// }

// func (pg *PostgresStore) DeleteTaskTemplate(id uuid.UUID) error {
// 	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, taskTemplateTableName)
// 	result, err := pg.db.Exec(query, title)
// 	if err != nil {
// 		return handleDbError(err, "task-template")
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return handleDbError(err, "task-template")
// 	}
// 	if rowsAffected == 0 {
// 		return errors.Wrapf(ae.ErrNotFound, "task template with title '%s' not found", title)
// 	}
// 	return nil
// }

// func (pg *PostgresStore) GetTaskTemplate(id uuid.UUID) (tp.TaskTemplate, error) {
// 	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s WHERE title = $1`, taskTemplateTableName)
// 	row := pg.db.QueryRow(query, title)

// 	var taskTemplate tp.TaskTemplate
// 	err := row.Scan(&taskTemplate.Id, &taskTemplate.Title, &taskTemplate.Description, &taskTemplate.Type)
// 	if err != nil {
// 		return tp.TaskTemplate{}, handleDbError(err, "task-template")
// 	}

// 	return taskTemplate, nil
// }

// func (pg *PostgresStore) ListTaskTemplates() ([]tp.TaskTemplate, error) {
// 	query := fmt.Sprintf(`SELECT id, title, description, type FROM %s`, taskTemplateTableName)
// 	rows, err := pg.db.Query(query)
// 	if err != nil {
// 		return nil, handleDbError(err, "task-template")
// 	}
// 	defer rows.Close()

// 	var taskTemplates = []tp.TaskTemplate{}
// 	for rows.Next() {
// 		var taskTemplate tp.TaskTemplate
// 		err = rows.Scan(&taskTemplate.Id, &taskTemplate.Title, &taskTemplate.Description, &taskTemplate.Type)
// 		if err != nil {
// 			return nil, handleDbError(err, "task-template")
// 		}
// 		taskTemplates = append(taskTemplates, taskTemplate)
// 	}

// 	return taskTemplates, nil
// }

// func (pg *PostgresStore) UpdateTaskTemplate(tt tp.TaskTemplate) (tp.TaskTemplate, error) {
// 	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2, type = $3 WHERE title = $4 returning id`, taskTemplateTableName)
// 	err := pg.db.QueryRow(query, taskTemplate.Title, taskTemplate.Description, taskTemplate.Type, title).Scan(&taskTemplate.Id)
// 	if err != nil {
// 		return tp.TaskTemplate{}, handleDbError(err, "task-template")
// 	}

// 	return taskTemplate, nil
// }
