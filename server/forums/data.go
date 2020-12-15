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
	//rows, err := s.Db.Query("SELECT id, nickname, interests FROM channels LIMIT 200")
		rows, err := s.Db.Query("SELECT f.id, f.name, f.topicKeyword FROM forums_sys.forums f")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
//		var users []string
		if err := rows.Scan(&f.ID, &f.Name, &f.TopicKeyword); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}

	//for _, forum := range res {
	//	id := forum.ID

/*		id := 1;
		
		usres, _ := s.Db.Query("SELECT u.nickname FROM forums_sys.interests i, forums_sys.users u WHERE i.forums_id=$1", id)
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
        res[1].Users = users
 
    }
    
*/

	if res == nil {
		res = make([]*Forum, 0)
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
	fmt.Println("received post request with nickname $1", Nickname)
	/*
	_, err := s.Db.Exec("INSERT INTO channels (nickname, interests) VALUES ($1)", Nickname, Interests)
	return err
	*/
	_, err := s.Db.Exec("INSERT INTO users (nickname) VALUES ($1)", Nickname)
	/*for i := 0; i< len(Interests); i++{
		_, err := s.Db.Exec("insert into forums_sys.interests (`users_id`, `forums_id`) values (
							(select u.id from forums_sys.users u where u.nickname = $1),
							(select c.id from forums_sys.forums c where c.topicKeyword = $2))", Nickname, i)
		}*/
	return err
}
