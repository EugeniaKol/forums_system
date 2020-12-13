package forums

import (
	"database/sql"
	"fmt"
)

// User defines user struct
type User struct {
	ID        int64    `json:"id"`
	Nickname  string   `json:"nickname"`
	Interests []string `json:"interests"`
}

// Store points to db
type Store struct {
	Db *sql.DB
}

// NewStore creates Store
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// ListForums recieves forums from db and lists them
func (s *Store) ListForums() ([]*User, error) {
	rows, err := s.Db.Query("SELECT id, nickname, interests FROM channels LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*User
	for rows.Next() {
		var c User
		if err := rows.Scan(&c.ID, &c.Nickname); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*User, 0)
	}
	return res, nil
}

// CreateUser creates user 
func (s *Store) CreateUser(Nickname string, Interests []string) error {
	if len(Nickname) < 0 {
		return fmt.Errorf("User nickname is not provided")
	}
	if len(Interests) == 0 {
		return fmt.Errorf("User interests is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO channels (nickname, interests) VALUES ($1)", Nickname, Interests)
	return err
}
