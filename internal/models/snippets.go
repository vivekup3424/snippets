package models

import (
	"database/sql"
	"log"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	query := `
        INSERT INTO snippets (title, content,created, expires)
        VALUES ($1, $2,current_timestamp, current_timestamp + interval '1 day' * $3)
        RETURNING id`

	var lastInsertId int
	err := m.DB.QueryRow(query, title, content, expires).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	query := `SELECT id,title,content, created, expires 
			FROM snippets
			WHERE id = $1
	`
	row := m.DB.QueryRow(query, id)
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
