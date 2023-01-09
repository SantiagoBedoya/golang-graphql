package users

import (
	"database/sql"
	"log"

	database "github.com/SantiagoBedoya/hackernews/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Create() {
	stmt, err := database.DB.Prepare("INSERT INTO users(username, password) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	hash, _ := hashPassword(u.Password)
	_, err = stmt.Exec(u.Username, hash)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *User) Authenticate() bool {
	stmt, err := database.DB.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Username)

	var hashPassword string
	err = row.Scan(&hashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return checkPassword(u.Password, hashPassword)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.DB.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)
	var id int
	err = row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return id, nil
}
