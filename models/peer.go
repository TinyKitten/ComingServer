package models

// PeerList すべてのピア
func PeerList(db XODB, offset, limit int) ([]*Peer, error) {
	sqlstr := `SELECT id, code, created_at, updated_At
	FROM peers
	LIMIT ?
	OFFSET ?`

	XOLog(sqlstr, limit, offset)
	q, err := db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*Peer{}
	for q.Next() {
		r := Peer{}
		err = q.Scan(
			&r.ID,
			&r.Code,
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

// PeerByToken トークンでピアを探す
func PeerByToken(db XODB, token string) (*Peer, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, code, token, created_at, updated_at ` +
		`FROM comingserver.peers ` +
		`WHERE token = ?`

	// run query
	XOLog(sqlstr, token)
	p := Peer{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, token).Scan(&p.ID, &p.Code, &p.Token, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
