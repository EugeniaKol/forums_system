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
	rows, err := s.Db.Query("SELECT f.id, f.name, f.topicKeyword FROM forums_sys.forums f")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.ID, &f.Name, &f.TopicKeyword); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}


	for _, forum := range res{
		id := forum.ID
		usres, _ := s.Db.Query("SELECT nickname FROM interests i, users u WHERE i.users_id=u.id AND i.forums_id= ?", id)
		if err != nil {
			return nil, err
		}
		defer usres.Close()
		var Nickname string
		var users []string
    	
    	for usres.Next() {

        	if err := usres.Scan(&Nickname); err != nil {
         	return nil, err
        }
        
        users = append(users, Nickname)
        forum.Users = users
    	}
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil 
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

