package engine

import (
	"database/sql"
	"path/filepath"
	"testing"
)

func TestDevinMigrationPreservesExistingHistoryAndSessionMapping(t *testing.T) {
	path := filepath.Join(t.TempDir(), "history.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// Model the previous installed schema, not only an empty database.
	for _, m := range migrations {
		if m.version >= 30 {
			break
		}
		tx, e := db.Begin()
		if e != nil {
			t.Fatal(e)
		}
		if e = m.apply(tx); e != nil {
			tx.Rollback()
			t.Fatal(e)
		}
		if e = tx.Commit(); e != nil {
			t.Fatal(e)
		}
	}
	if _, err = db.Exec(`PRAGMA user_version=29; INSERT INTO sessions(id,project_key,title,created_at,updated_at) VALUES('chat','project','keep me','now','now')`); err != nil {
		t.Fatal(err)
	}
	if err = migrate(db); err != nil {
		t.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO external_agent_sessions VALUES('chat','devin','project','official-session')`); err != nil {
		t.Fatal(err)
	}
	db.Close()
	db, err = sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err = migrate(db); err != nil {
		t.Fatal(err)
	}
	var title, id string
	if err = db.QueryRow(`SELECT title FROM sessions WHERE id='chat'`).Scan(&title); err != nil {
		t.Fatal(err)
	}
	if err = db.QueryRow(`SELECT external_id FROM external_agent_sessions WHERE session_id='chat' AND provider='devin' AND project_key='project'`).Scan(&id); err != nil {
		t.Fatal(err)
	}
	if title != "keep me" || id != "official-session" {
		t.Fatalf("history=%q session=%q", title, id)
	}
}
