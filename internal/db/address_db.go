package db

func (db *Postgres) DeleteAddress(id int, userId int) error {
	stmt, err := db.Db.Prepare(`DELETE FROM address WHERE id=$1 AND user_id=$2`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, userId)
	if err != nil {
		return err
	}
	return nil

}
