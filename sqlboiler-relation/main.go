package main

import (
	"context"
	"database/sql"
	"fmt"

	"./app/db"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	// Connect To DB
	con, err := sql.Open("postgres", "host=localhost port=5432 user=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	boil.SetDB(con)

	/*
	// CREATE
	user := db.User{
		Name: null.StringFrom("hoge"),
		ID: 3,
	}
	c := db.CreditCard{
		Number: null.StringFrom("n"),
		ID: 21,
	}
	c2 := db.CreditCard{
		Number: null.StringFrom("n2"),
		ID: 22,
	}
	user.InsertGP(context.Background(), boil.Infer())

	user.SetCreditCardsGP(context.Background(),true,&c,&c2)
	*/

	// READ
	users := db.Users().OneGP(context.Background())

	cards := users.CreditCards().AllGP(context.Background())

	for _, c := range cards {
		fmt.Println(c.Number)
	}

}