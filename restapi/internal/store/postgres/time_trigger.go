package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var timeTriggerTableName = "timetrigger"

func (s *Store) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	tt.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, quantity, timeunit_title, task_id) VALUES ($1, $2, $3, $4)`, timeTriggerTableName)
	_, err := s.db.Exec(query, tt.Id, tt.Quantity, tt.TimeUnit, tt.TaskId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	return tt, nil
}

func (s *Store) DeleteTimeTrigger(ttId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, timeTriggerTableName)
	_, err := s.db.Exec(query, ttId)
	if err != nil {
		return handleDbError(err, "time-trigger")
	}

	return nil
}

func (s *Store) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_title, task_id FROM %s`, timeTriggerTableName)
	rows, err := s.db.Query(query)
	if err != nil {
		return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}
	defer rows.Close()

	var ttgs []tp.TimeTrigger
	for rows.Next() {
		var tt tp.TimeTrigger
		err := rows.Scan(&tt.Id, &tt.Quantity, &tt.TimeUnit, &tt.TaskId)
		if err != nil {
			return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
		}
		ttgs = append(ttgs, tt)
	}

	return ttgs, nil
}

func (s *Store) GetTimeTrigger(ttId uuid.UUID) (tp.TimeTrigger, error) {
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_title, task_id FROM %s WHERE id=$1`, timeTriggerTableName)
	row := s.db.QueryRow(query, ttId)

	var tt tp.TimeTrigger
	err := row.Scan(&tt.Id, &tt.Quantity, &tt.TimeUnit, &tt.TaskId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	return tt, nil
}

func (s *Store) UpdateTimeTrigger(ttId uuid.UUID, tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	query := fmt.Sprintf(`UPDATE %s SET quantity=$1, timeunit_title=$2, task_id=$3 WHERE id=$4`, timeTriggerTableName)
	_, err := s.db.Exec(query, tt.Quantity, tt.TimeUnit, tt.TaskId, ttId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	return tt, nil
}
