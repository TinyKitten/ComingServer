package models

// PodList すべてのポッド
func PodList(db XODB, offset, limit int) ([]*Pod, error) {
	sqlstr := `SELECT id, code, latitude, longitude, approaching, created_at, updated_at
	FROM pods
	LIMIT ?
	OFFSET ?`

	XOLog(sqlstr, limit, offset)
	q, err := db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*Pod{}
	for q.Next() {
		r := Pod{}
		err = q.Scan(
			&r.ID,
			&r.Code,
			&r.Latitude,
			&r.Longitude,
			&r.Approaching,
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

// PodByToken トークンでポッドを探す
func PodByToken(db XODB, token string) (*Pod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, code, latitude, longitude, approaching, token, created_at, updated_at ` +
		`FROM comingserver.pods ` +
		`WHERE token = ?`

	// run query
	XOLog(sqlstr, token)
	p := Pod{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, token).Scan(&p.ID, &p.Code, &p.Latitude, &p.Longitude, &p.Approaching, &p.Token, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
