package equipment

import d "github.com/jocarde/personacmms/pkg/db"

var db = d.Db

type Equipment struct {
	id          int
	title       string
	description string
}

// Create creates a new equipment record in the database
func Create(title string, description string) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO equipment (title, description) VALUES ($1, $2) RETURNING id`
	res, err := db.Exec(query, title, description)
	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()
	return int(id), nil
}

// Update updates an equipment record in the database
func Update(id int, title string, description string) error {
	query := `UPDATE equipment SET title = $1, description = $2 WHERE id = $3`
	_, err := db.Exec(query, title, description, id)

	return err
}

// Get returns an equipment record from the database
func Get(id int) (Equipment, error) {
	var e Equipment
	query := `SELECT id, title, description FROM equipment WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&e.id, &e.title, &e.description)
	return e, err
}

// GetAll returns all equipment records from the database
func GetAll() ([]Equipment, error) {
	var equipment []Equipment
	query := `SELECT id, title, description FROM equipment`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var e Equipment
		err = rows.Scan(&e.id, &e.title, &e.description)
		if err != nil {
			panic(err)
		}
		equipment = append(equipment, e)
	}

	return equipment, nil
}

func Delete(id int) error {
	query := `DELETE FROM equipment WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		panic(err)
	}

	return nil
}
