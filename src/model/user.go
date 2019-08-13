package model

import (
	"fmt"
	// "net/http"
	// "strconv"

	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	id   int
	name string
}

var db *sql.DB

func FindAll() ([]User, error) {
	rows, err := db.Query("select * from users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    users := []User{}
     
    for rows.Next(){
        p := User{}
        err := rows.Scan(&p.id, &p.name)
        if err != nil{
            fmt.Println(err)
            continue
        }
        users = append(users, p)
	}
	return users, nil
    // for _, p := range products{
    //     fmt.Println(p.id, p.model, p.company, p.price)
    // }
}

func Create(name string) (*User, error) {
	sqlStatement := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		return nil, err
	}
	fmt.Println("New record ID is: ", id)

	return &User{
		id:   id,
		name: name,
	}, nil
}

func InitDb() {
	// config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", "5432",
		"api", "api", "test-go")

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB server successfully connected!")
}

func CloseDb() {
	db.Close()
}
