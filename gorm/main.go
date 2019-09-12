package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// User has belong to Star
type User struct {
	Name        string
	ID          uint
	CreditCards []CreditCard
}

type CreditCard struct {
	UserID uint
	Number string
	ID     uint
}

// Array of User
type Users []User

func main() {
	// Connect to DB
	db, err := gorm.Open("postgres", "host=localhost port=15432 user=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &CreditCard{})

	// CREATE
	db.Create(&User{Name: "hoge",
		ID: 20,
		CreditCards: []CreditCard{
			{
				Number: "1x",
				ID:     1,
			},
			{
				Number: "2x",
				ID:     2,
			},
		},
	})

	u := User{}
	u.ID = 20
	cs := []CreditCard{}
	db.Model(&u).Related(&cs)
	fmt.Println(cs)
	// READ
	users := Users{}
	db.Find(&users) // SELECT * FROM users

	user := User{}
	db.Take(&user) // SELECT * FROM users LIMIT 1;

	// UPDATE
	db.Model(&user).Update("Name", "huga")

	// DELETE
	db.Delete(&user)

}
