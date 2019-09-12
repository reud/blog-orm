package main

import (
	"fmt"

	model "./models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func main() {
	// TODO: portを15432に直す
	engine, err := xorm.NewEngine("postgres", "dbname=postgres host=localhost port=5432 user=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer engine.Close()

	user := model.Users{Id: 3, Name: "hoge"}

	// CREATE
	engine.Insert(&user)

	var users []model.Users

	// READ
	engine.Find(&users)

	// UPDATE
	user.Name = "huga"
	engine.ID(3).Update(&user)

	// DELETE
	engine.ID(3).Delete(&user)

}
