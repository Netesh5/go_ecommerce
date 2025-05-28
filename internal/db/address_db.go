package db

func (db *Postgres) DeleteAddress(id int) error {
	stmt, err := db.Db.Prepare(`DELETE FROM address WHERE id=$1`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil

}
