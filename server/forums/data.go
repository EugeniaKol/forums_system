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

type Interest struct {
	Nickname   string   `json:"nickname"`
	Forums_id  int64	`json:"forums_id"`
}

type Forum struct {
	ID        		int64    `json:"id"`
	Name  			string   `json:"name"`
	TopicKeyword  	string   `json:"topicKeyword"`
	Users 			[]string `json:"users"`
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
func (s *Store) ListForums() ([]*Forum, error) {
	frows, err := s.Db.Query("SELECT f.id, f.name, f.topicKeyword FROM forums_sys.forums f")

	if err != nil {
		return nil, err
	}

	defer frows.Close()

	var fres []*Forum
	for frows.Next() {
		var f Forum
		if err := frows.Scan(&f.ID, &f.Name, &f.TopicKeyword); err != nil {
			return nil, err
		}
		fres = append(fres, &f)
	}

////////////////////////////////

	irows, err := s.Db.Query("SELECT u.nickname, i.forums_id FROM interests i, users u WHERE i.users_id=u.id")

	if err != nil {
		return nil, err
	}

	defer irows.Close()

	var ires []*Interest
	for irows.Next() {
		var i Interest
		if err := irows.Scan(&i.Nickname, &i.Forums_id); err != nil {
			return nil, err
		}
		ires = append(ires, &i)
	}

	for _, forum := range fres{
		var users []string
		for _, interest := range ires{
			if forum.ID == interest.Forums_id{
				users = append(users, interest.Nickname)
			}
		}
		forum.Users = users
	}

	if fres == nil {
		fres = make([]*Forum, 0)
	}

	return fres, nil 
}

// CreateUser creates user 
func (s *Store) CreateUser(nickname string , Interests []string) error {
	if len(nickname) < 0 {
		return fmt.Errorf("User nickname is not provided")
	}
	statement1, err := s.Db.Prepare("INSERT INTO users (nickname) VALUES (?)")
				statement1.Exec(nickname);
	
	statement2, err := s.Db.Prepare("INSERT INTO interests (users_id, forums_id) VALUES ((select u.id from forums_sys.users u where u.nickname = ?),(select c.id from forums_sys.forums c where c.topicKeyword =?))")
				for _, i := range Interests {
				statement2.Exec(nickname, i)
				}

	return err
}

