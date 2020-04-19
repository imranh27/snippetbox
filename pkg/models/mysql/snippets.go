package mysql

import (
	"database/sql"
	"errors"
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

	//SQL statement for retrieving data.
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	//initialise a pointer to a new zeroed Snippet struct
	s := &models.Snippet{}

	//map query results to snippet struct
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

//Return the 10 most recently used snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := "SELECT id, title, content, created, expires " +
		"FROM snippets " +
		"WHERE expires > UTC_TIMESTAMP() " +
		"ORDER BY created DESC LIMIT 10"

	//Use the Query method on connection pool to execute statement
	//returns sql.Rows result set.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//initialise empty slice to hold models.snippets objects.
	var snippets []*models.Snippet

	//iterate through resultset
	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	//find any error through iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
