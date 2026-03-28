package db

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

var ErrNoTasks = errors.New("no tasks available")

type List struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Task struct {
	ID        int64   `json:"id"`
	ListID    int64   `json:"list_id"`
	Title     string  `json:"title"`
	Status    string  `json:"status"`
	Meta      *string `json:"meta,omitempty"`
	Result    *string `json:"result,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ListStatus struct {
	ListID  int64  `json:"list_id"`
	Name    string `json:"name"`
	Total   int    `json:"total"`
	Pending int    `json:"pending"`
	Active  int    `json:"active"`
	Done    int    `json:"done"`
	Failed  int    `json:"failed"`
}

const schema = `
CREATE TABLE IF NOT EXISTS lists (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	list_id INTEGER NOT NULL REFERENCES lists(id),
	title TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending','active','done','failed')),
	meta TEXT,
	result TEXT,
	reason TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_tasks_list_status ON tasks(list_id, status);
`

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(wal)&_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
