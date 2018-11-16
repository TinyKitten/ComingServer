package models

// UserList すべてのユーザ
func UserList(db XODB, offset, limit int) ([]*User, error) {
	sqlstr := `SELECT screen_name, privilege_id, created_at, updated_at
	FROM users
	LIMIT ?
	OFFSET ?`

	XOLog(sqlstr, limit, offset)
	q, err := db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*User{}
	for q.Next() {
		r := User{}
		err = q.Scan(
			&r.ID,
			&r.ScreenName,
			&r.PrivilegeID,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &r)
	}

	return res, nil
}
