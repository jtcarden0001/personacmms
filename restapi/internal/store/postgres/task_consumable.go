package postgres

import (
	"database/sql"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type PreventativeTaskConsumable interface {
	CreatePreventativeTaskConsumable(int, int, string) error
	DeletePreventativeTaskConsumable(int, int) error
	GetAllPreventativeTaskConsumable() ([]tp.PreventativeTaskConsumable, error)
	GetAllPreventativeTaskConsumableByPreventativeTaskId(int) ([]tp.PreventativeTaskConsumable, error)
	GetPreventativeTaskConsumable(int, int) (tp.PreventativeTaskConsumable, error)
	UpdatePreventativeTaskConsumable(int, int, string) error
}

func (pg *Store) CreatePreventativeTaskConsumable(preventativeTaskId int, consumableId int, quantity string) error {
	query := `INSERT INTO preventativeTask_consumable (preventativeTask_id, consumable_id, quantity_note) VALUES ($1, $2, $3)`
	_, err := pg.db.Exec(query, preventativeTaskId, consumableId, quantity)

	return err
}

func (pg *Store) DeletePreventativeTaskConsumable(preventativeTaskId int, consumableId int) error {
	query := `DELETE FROM preventativeTask_consumable WHERE preventativeTask_id = $1 AND consumable_id = $2`
	_, err := pg.db.Exec(query, preventativeTaskId, consumableId)

	return err
}

func (pg *Store) GetAllPreventativeTaskConsumable() ([]tp.PreventativeTaskConsumable, error) {
	query := `SELECT preventativeTask_id, consumable_id, quantity_note FROM preventativeTask_consumable`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populatePreventativeTaskConsumableList(rows)
}

func (pg *Store) GetAllPreventativeTaskConsumableByPreventativeTaskId(preventativeTaskId int) ([]tp.PreventativeTaskConsumable, error) {
	query := `SELECT preventativeTask_id, consumable_id, quantity_note FROM preventativeTask_consumable WHERE preventativeTask_id = $1`
	rows, err := pg.db.Query(query, preventativeTaskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populatePreventativeTaskConsumableList(rows)
}

func (pg *Store) GetPreventativeTaskConsumable(consumableId int, preventativeTaskId int) (tp.PreventativeTaskConsumable, error) {
	query := `SELECT preventativeTask_id, consumable_id, quantity_note FROM preventativeTask_consumable WHERE preventativeTask_id = $1 AND consumable_id = $2`
	var tc tp.PreventativeTaskConsumable
	err := pg.db.QueryRow(query, preventativeTaskId, consumableId).Scan(&tc.PreventativeTaskId, &tc.ConsumableId, &tc.QuantityNote)

	return tc, err
}

func (pg *Store) UpdatePreventativeTaskConsumable(preventativeTaskId int, consumableId int, quantity string) error {
	query := `UPDATE preventativeTask_consumable SET quantity_note = $1 WHERE preventativeTask_id = $2 AND consumable_id = $3`
	_, err := pg.db.Exec(query, quantity, preventativeTaskId, consumableId)

	return err
}

func populatePreventativeTaskConsumableList(rows *sql.Rows) ([]tp.PreventativeTaskConsumable, error) {
	var preventativeTaskConsumables []tp.PreventativeTaskConsumable
	for rows.Next() {
		var tc tp.PreventativeTaskConsumable
		err := rows.Scan(&tc.PreventativeTaskId, &tc.ConsumableId, &tc.QuantityNote)
		if err != nil {
			return nil, err
		}
		preventativeTaskConsumables = append(preventativeTaskConsumables, tc)
	}

	return preventativeTaskConsumables, nil
}
