package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTaskConsumableTable = "task_consumable"

func (pg *Store) CreateTaskConsumable(consumable tp.TaskConsumable) (tp.TaskConsumable, error) {
	query := fmt.Sprintf("INSERT INTO %s (task_id, consumable_id, quantity_note) VALUES ($1, $2, $3)", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, consumable.TaskId, consumable.ConsumableId, consumable.QuantityNote)
	if err != nil {
		return tp.TaskConsumable{}, handleDbError(err, "task-consumable")
	}

	return consumable, nil
}

func (pg *Store) DeleteTaskConsumable(atId, cId tp.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE task_id = $1 AND consumable_id = $2", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, atId, cId)
	if err != nil {
		return handleDbError(err, "task-consumable")
	}

	return nil
}

func (pg *Store) ListTaskConsumables() ([]tp.TaskConsumable, error) {
	query := fmt.Sprintf("SELECT task_id, consumable_id, quantity_note FROM %s", assetTaskConsumableTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "task-consumable")
	}
	defer rows.Close()

	var at []tp.TaskConsumable
	for rows.Next() {
		var e tp.TaskConsumable
		err = rows.Scan(&e.TaskId, &e.ConsumableId, &e.QuantityNote)
		if err != nil {
			return nil, handleDbError(err, "task-consumable")
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *Store) GetTaskConsumable(atId, cId tp.UUID) (tp.TaskConsumable, error) {
	query := fmt.Sprintf("SELECT task_id, consumable_id, quantity_note FROM %s WHERE task_id = $1 AND consumable_id = $2", assetTaskConsumableTable)
	var e tp.TaskConsumable
	err := pg.db.QueryRow(query, atId, cId).Scan(&e.TaskId, &e.ConsumableId, &e.QuantityNote)
	if err != nil {
		return tp.TaskConsumable{}, handleDbError(err, "task-consumable")
	}

	return e, nil
}

func (pg *Store) UpdateTaskConsumable(atc tp.TaskConsumable) (tp.TaskConsumable, error) {
	query := fmt.Sprintf("UPDATE %s SET quantity_note = $1 WHERE task_id = $2 AND consumable_id = $3", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, atc.QuantityNote, atc.TaskId, atc.ConsumableId)
	if err != nil {
		return tp.TaskConsumable{}, handleDbError(err, "task-consumable")
	}

	return atc, nil
}
