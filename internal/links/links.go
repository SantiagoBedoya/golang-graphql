package links

import (
	"log"

	database "github.com/SantiagoBedoya/hackernews/internal/pkg/db/mysql"
	"github.com/SantiagoBedoya/hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func GetAll() []Link {
	stmt, err := database.DB.Prepare("SELECT L.id, L.title, L.address, L.user_id, U.username from links L INNER JOIN users U on L.user_id = U.id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}

func (l Link) Save() int64 {
	stmt, err := database.DB.Prepare("INSERT INTO links(title, address, user_id) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(l.Title, l.Address, l.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	log.Print("Row inserted")
	return id
}
