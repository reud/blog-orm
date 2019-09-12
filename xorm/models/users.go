package models

type Users struct {
	Name string `xorm:"TEXT"`
	Id   int    `xorm:"not null pk autoincr INTEGER"`
}
