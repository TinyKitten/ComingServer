package models

import "errors"

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

// PodUpdateToken ポッドのトークンを更新
func PodUpdateToken(db XODB, podID uint64, token string) error {
	const sqlstr = `UPDATE pods
	SET token = ?,
	updated_at = CURRENT_TIMESTAMP
	WHERE id=?`

	XOLog(sqlstr, &token, podID)
	_, err := db.Exec(sqlstr, token, podID)
	return err
}

// PodUpdateInput ポッドアップデート入力
type PodUpdateInput struct {
	ID        *uint64
	Latitude  *float64
	Longitude *float64
	Code      *string
}

// PodUpdate ポッドの情報を更新
func PodUpdate(db XODB, input *PodUpdateInput) error {
	if input.ID == nil {
		return errors.New("ID is required")
	}
	existsPod, err := PodByID(db, *input.ID)
	if err != nil {
		return err
	}

	const sqlstr = `UPDATE pods
	SET latitude = ?,
	longitude = ?,
	code = ?,
	updated_at = CURRENT_TIMESTAMP
	WHERE id=?`

	if input.Latitude == nil {
		input.Latitude = &existsPod.Latitude
	}
	if input.Longitude == nil {
		input.Longitude = &existsPod.Longitude
	}
	if input.Code == nil {
		input.Code = &existsPod.Code
	}

	XOLog(sqlstr, &input.Latitude, &input.Longitude, &input.Code, &input.ID)
	_, err = db.Exec(sqlstr, &input.Latitude, &input.Longitude, &input.Code, &input.ID)

	return err
}
