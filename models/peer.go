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
