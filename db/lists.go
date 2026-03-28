package db

import "database/sql"

func CreateList(db *sql.DB, name string) (*List, error) {
	l := &List{}
	err := db.QueryRow(
		`INSERT INTO lists (name) VALUES (?) RETURNING id, name, created_at, updated_at`,
		name,
	).Scan(&l.ID, &l.Name, &l.CreatedAt, &l.UpdatedAt)
	return l, err
}

func GetAllLists(db *sql.DB) ([]ListStatus, error) {
	rows, err := db.Query(`
		SELECT l.id, l.name,
			COUNT(t.id) as total,
			SUM(CASE WHEN t.status = 'pending' THEN 1 ELSE 0 END) as pending,
			SUM(CASE WHEN t.status = 'active' THEN 1 ELSE 0 END) as active,
			SUM(CASE WHEN t.status = 'done' THEN 1 ELSE 0 END) as done,
			SUM(CASE WHEN t.status = 'failed' THEN 1 ELSE 0 END) as failed
		FROM lists l
		LEFT JOIN tasks t ON l.id = t.list_id
		GROUP BY l.id
		ORDER BY l.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []ListStatus
	for rows.Next() {
		var s ListStatus
		if err := rows.Scan(&s.ListID, &s.Name, &s.Total, &s.Pending, &s.Active, &s.Done, &s.Failed); err != nil {
			return nil, err
		}
		lists = append(lists, s)
	}
	return lists, rows.Err()
}
