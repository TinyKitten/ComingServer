package models

// PodList すべてのポッド
func PodList(db XODB, offset, limit int) ([]*Pod, error) {
	sqlstr := `SELECT id, code, latitude, longitude, rumbling, created_at, updated_At
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
			&r.Rumbling,
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
