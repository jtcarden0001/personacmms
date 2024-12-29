package postgres

import (
	"fmt"

	"github.com/google/uuid"
	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var consumableTableName = "consumable"
var tConsumableTableName = "task_consumable"
var woConsumableTableName = "workorder_consumable"

func (pg *PostgresStore) AssociateConsumableWithTask(taskId uuid.UUID, consumableId uuid.UUID, qNote string) (tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (task_id, consumable_id, quantity_note) 
			VALUES ($1, $2, $3)`,
		tConsumableTableName)

	_, err := pg.db.Exec(query, consumableId, taskId, qNote)
	if err != nil {
		return tp.ConsumableQuantity{}, handleDbError(err, "consumable")
	}

	return pg.GetTaskConsumableQuantity(taskId, consumableId)
}

func (pg *PostgresStore) AssociateConsumableWithWorkOrder(workOrderId uuid.UUID, consumableId uuid.UUID, qNote string) (tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (work_order_id, consumable_id, quantity_note) 
			VALUES ($1, $2, $3)`,
		woConsumableTableName)

	_, err := pg.db.Exec(query, workOrderId, consumableId, qNote)
	if err != nil {
		return tp.ConsumableQuantity{}, handleDbError(err, "consumable")
	}

	return pg.GetWorkOrderConsumableQuantity(workOrderId, consumableId)
}

func (pg *PostgresStore) CreateConsumable(c tp.Consumable) (tp.Consumable, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, title) 
			VALUES ($1, $2)`,
		consumableTableName)

	_, err := pg.db.Exec(query, c.Id, c.Title)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) DeleteConsumable(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		consumableTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "consumable")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "consumable")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "consumable with id %s not found", id.String())
	}
	return nil
}

func (pg *PostgresStore) DisassociateConsumableWithTask(taskId uid.UUID, consumableId uid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE task_id = $1 AND consumable_id = $2`,
		tConsumableTableName)

	result, err := pg.db.Exec(query, taskId, consumableId)
	if err != nil {
		return handleDbError(err, "consumable")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "consumable")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "consumable id:[%s] not associated with task id:[%s]", consumableId, taskId)
	}

	return nil
}

func (pg *PostgresStore) DisassociateConsumableWithWorkOrder(workOrderId uid.UUID, consumableId uid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE work_order_id = $1 AND consumable_id = $2`,
		woConsumableTableName)

	result, err := pg.db.Exec(query, workOrderId, consumableId)
	if err != nil {
		return handleDbError(err, "consumable")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "consumable")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "consumable id:[%s] not associated with work order id:[%s]", consumableId, workOrderId)
	}

	return nil
}

func (pg *PostgresStore) GetConsumable(id uid.UUID) (tp.Consumable, error) {
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s 
			WHERE id = $1`,
		consumableTableName)

	var c tp.Consumable
	err := pg.db.QueryRow(query, id).Scan(&c.Id, &c.Title)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) GetTaskConsumableQuantity(taskId uid.UUID, consumableId uid.UUID) (tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			SELECT c.id, c.title, tc.quantity_note 
			FROM %s c JOIN %s tc ON c.id = tc.consumable_id 
			WHERE tc.task_id = $1 AND tc.consumable_id = $2`,
		consumableTableName, tConsumableTableName)

	var c tp.ConsumableQuantity
	err := pg.db.QueryRow(query, taskId, consumableId).Scan(&c.Id, &c.Title, &c.Quantity)
	if err != nil {
		return tp.ConsumableQuantity{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) GetWorkOrderConsumableQuantity(woId uid.UUID, consumableId uid.UUID) (tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			SELECT c.id, c.title, tc.quantity_note 
			FROM %s c JOIN %s tc ON c.id = tc.consumable_id 
			WHERE tc.workorder_id = $1 AND tc.consumable_id = $2`,
		consumableTableName, woConsumableTableName)

	var c tp.ConsumableQuantity
	err := pg.db.QueryRow(query, woId, consumableId).Scan(&c.Id, &c.Title, &c.Quantity)
	if err != nil {
		return tp.ConsumableQuantity{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) ListConsumables() ([]tp.Consumable, error) {
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s`,
		consumableTableName)

	var consumables = []tp.Consumable{}
	rows, err := pg.db.Query(query)
	if err != nil {
		return consumables, handleDbError(err, "consumable")
	}
	defer rows.Close()

	for rows.Next() {
		var c tp.Consumable
		err = rows.Scan(&c.Id, &c.Title)
		if err != nil {
			return nil, handleDbError(err, "consumable")
		}
		consumables = append(consumables, c)
	}

	return consumables, nil
}

func (pg *PostgresStore) ListConsumablesByTask(taskId uid.UUID) ([]tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			SELECT c.id, c.title, tc.quantity_note 
			FROM %s c JOIN %s tc ON c.id = tc.consumable_id 
			WHERE tc.task_id = $1`,
		consumableTableName, tConsumableTableName)

	var consumables = []tp.ConsumableQuantity{}
	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return consumables, handleDbError(err, "consumable")
	}
	defer rows.Close()

	for rows.Next() {
		var c tp.ConsumableQuantity
		err = rows.Scan(&c.Id, &c.Title, &c.Quantity)
		if err != nil {
			return nil, handleDbError(err, "consumable")
		}
		consumables = append(consumables, c)
	}

	return consumables, nil
}

func (pg *PostgresStore) ListConsumablesByWorkOrder(workOrderId uid.UUID) ([]tp.ConsumableQuantity, error) {
	query := fmt.Sprintf(`
			SELECT c.id, c.title, wc.quantity_note
			FROM %s c JOIN %s wc ON c.id = wc.consumable_id
			WHERE wc.work_order_id = $1`,
		consumableTableName, woConsumableTableName)

	var consumables = []tp.ConsumableQuantity{}
	rows, err := pg.db.Query(query, workOrderId)
	if err != nil {
		return consumables, handleDbError(err, "consumable")
	}
	defer rows.Close()

	for rows.Next() {
		var c tp.ConsumableQuantity
		err = rows.Scan(&c.Id, &c.Title, &c.Quantity)
		if err != nil {
			return nil, handleDbError(err, "consumable")
		}
		consumables = append(consumables, c)
	}

	return consumables, nil
}

func (pg *PostgresStore) UpdateConsumable(c tp.Consumable) (tp.Consumable, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET title = $1 
			WHERE id = $2`,
		consumableTableName)

	result, err := pg.db.Exec(query, c.Title, c.Id)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	if rowsAffected == 0 {
		return tp.Consumable{}, errors.Wrapf(ae.ErrNotFound, "consumable with id %s not found", c.Id)
	}

	return c, nil
}
