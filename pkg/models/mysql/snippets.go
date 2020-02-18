package mysql

import (
	"database/sql"
	"github.com/imranh27/snippetbox/pkg/models"
)

//Define snippet model that wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

//Insert a new snippet in to the DB.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := "INSERT INTO snippets(title, content, created, expires)" +
				" VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY) )"

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

//Return specific snippet based on ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

//Return the 10 most recently used snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
