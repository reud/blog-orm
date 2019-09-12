package main

import (
	"context"
	"database/sql"
	"fmt"

	"./app/db"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	// Connect To DB
	con, err := sql.Open("postgres", "host=localhost port=15432 user=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	boil.SetDB(con)
	// 発行されたSQL文が見れるようになる。
	boil.DebugMode = true

	// CREATE
	user := db.User{Name: null.StringFrom("hoge"), ID: 1}
	user.InsertGP(context.Background(), boil.Infer())

	// READ
	users := db.Users().AllGP(context.Background())

	fmt.Println(users)

	user = db.User{
		ID:   1,
		Name: null.StringFrom("huga"),
	}
	// UPDATE
	user.UpdateGP(context.Background(), boil.Infer())

	// DELETE
	user.DeleteGP(context.Background())

}


