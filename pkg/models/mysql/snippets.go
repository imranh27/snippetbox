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
	return 0, nil
}

//Return specific snippet based on ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

//Return the 10 most recently used snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
