////package main
////
////import (
////	"fmt"
////	"github.com/gorilla/mux"
////	"net/http"
////)
////
////func main() {
////
////	r := mux.NewRouter()
////
////	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
////		vars := mux.Vars(r)
////		title := vars["title"]
////		page := vars["page"]
////
////		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
////	})
////
////	fs := http.FileServer(http.Dir("static/"))
////	http.Handle("/static/", http.StripPrefix("/static/", fs))
////
////	http.ListenAndServe(":8080", r)
////}
//
//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := db.Ping(); err != nil {
//		log.Fatal(err)
//	}
//
//	{ // Create a new table
//		query := `
//            CREATE TABLE users (
//                id INT AUTO_INCREMENT,
//                username TEXT NOT NULL,
//                password TEXT NOT NULL,
//                created_at DATETIME,
//                PRIMARY KEY (id)
//            );`
//
//		if _, err := db.Exec(query); err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	{ // Insert a new user
//		username := "johndoe"
//		password := "secret"
//		createdAt := time.Now()
//
//		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		id, err := result.LastInsertId()
//		fmt.Println(id)
//	}
//
//	{ // Query a single user
//		var (
//			id        int
//			username  string
//			password  string
//			createdAt time.Time
//		)
//
//		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
//		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(id, username, password, createdAt)
//	}
//
//	{ // Query all users
//		type user struct {
//			id        int
//			username  string
//			password  string
//			createdAt time.Time
//		}
//
//		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer rows.Close()
//
//		var users []user
//		for rows.Next() {
//			var u user
//
//			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
//			if err != nil {
//				log.Fatal(err)
//			}
//			users = append(users, u)
//		}
//		if err := rows.Err(); err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("%#v", users)
//	}
//
//	{
//		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	deck := []string{"spade", "heart", "club", "diamond"}
	cards := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	randIndex := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch d := deck[randIndex.Intn(4)]; d {
	case "spade":
		fmt.Println(cards[randIndex.Intn(13)], "SPADE")
	case "club":
		fmt.Println(cards[randIndex.Intn(13)], "CLUB")
	case "heart":
		fmt.Println(cards[randIndex.Intn(13)], "HEART")
	case "diamond":
		fmt.Println(cards[randIndex.Intn(13)], "DIAMOND")
	default:
		fmt.Println("OOPS!!")
	}
}
