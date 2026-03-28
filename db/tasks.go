package db

import (
	"database/sql"
	"fmt"
)

func scanTask(row interface{ Scan(...any) error }) (*Task, error) {
	t := &Task{}
	err := row.Scan(&t.ID, &t.ListID, &t.Title, &t.Status, &t.Meta, &t.Result, &t.Reason, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

const taskCols = `id, list_id, title, status, meta, result, reason, created_at, updated_at`

func AddTask(db *sql.DB, listID int64, title string, meta string) (*Task, error) {
	var metaPtr *string
	if meta != "" {
		metaPtr = &meta
	}
	return scanTask(db.QueryRow(
		`INSERT INTO tasks (list_id, title, meta) VALUES (?, ?, ?) RETURNING `+taskCols,
		listID, title, metaPtr,
	))
}

func GetTask(db *sql.DB, taskID int64) (*Task, error) {
	t, err := scanTask(db.QueryRow(`SELECT `+taskCols+` FROM tasks WHERE id = ?`, taskID))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %d not found", taskID)
	}
	return t, err
}

func ClaimNextTask(db *sql.DB, listID int64) (*Task, error) {
	t, err := scanTask(db.QueryRow(`
		UPDATE tasks SET status = 'active', updated_at = CURRENT_TIMESTAMP
		WHERE id = (SELECT id FROM tasks WHERE list_id = ? AND status = 'pending' ORDER BY id LIMIT 1)
		RETURNING `+taskCols,
		listID,
	))
	if err == sql.ErrNoRows {
		return nil, ErrNoTasks
	}
	return t, err
}

func CompleteTask(db *sql.DB, taskID int64, result string) (*Task, error) {
	var resultPtr *string
	if result != "" {
		resultPtr = &result
	}
	t, err := scanTask(db.QueryRow(
		`UPDATE tasks SET status = 'done', result = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND status = 'active'
		RETURNING `+taskCols,
		resultPtr, taskID,
	))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %d not found or not in active state", taskID)
	}
	return t, err
}

func FailTask(db *sql.DB, taskID int64, reason string) (*Task, error) {
	var reasonPtr *string
	if reason != "" {
		reasonPtr = &reason
	}
	t, err := scanTask(db.QueryRow(
		`UPDATE tasks SET status = 'failed', reason = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND status = 'active'
		RETURNING `+taskCols,
		reasonPtr, taskID,
	))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %d not found or not in active state", taskID)
	}
	return t, err
}

func ResetTask(db *sql.DB, taskID int64) (*Task, error) {
	t, err := scanTask(db.QueryRow(
		`UPDATE tasks SET status = 'pending', result = NULL, reason = NULL, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND status IN ('active', 'failed')
		RETURNING `+taskCols,
		taskID,
	))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %d not found or not in active/failed state", taskID)
	}
	return t, err
}

func GetListStatus(db *sql.DB, listID int64) (*ListStatus, error) {
	s := &ListStatus{}
	err := db.QueryRow(`
		SELECT l.id, l.name,
			COUNT(t.id) as total,
			SUM(CASE WHEN t.status = 'pending' THEN 1 ELSE 0 END) as pending,
			SUM(CASE WHEN t.status = 'active' THEN 1 ELSE 0 END) as active,
			SUM(CASE WHEN t.status = 'done' THEN 1 ELSE 0 END) as done,
			SUM(CASE WHEN t.status = 'failed' THEN 1 ELSE 0 END) as failed
		FROM lists l
		LEFT JOIN tasks t ON l.id = t.list_id
		WHERE l.id = ?
		GROUP BY l.id`,
		listID,
	).Scan(&s.ListID, &s.Name, &s.Total, &s.Pending, &s.Active, &s.Done, &s.Failed)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("list %d not found", listID)
	}
	return s, err
}
