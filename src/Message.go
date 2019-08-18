package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Message struct {
	gorm.Model
	ID		  	string 		`json:"id"`
	Subject  	string 		`json:"subject"`
	User     	string 		`json:"user"`
	Message  	string 		`json:"message"`
	To			int 		`json:"to"`
	From 		int 		`json:"from"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}

func (Message) findFirst(id int) Message {
	var self Message
	getDB().First(&self, id)
	return self
}

//func (Message) findOneBy(criteria map[string]string) Message {
// Message.findOneBy("user = ?", "6")
func (Message) findOneBy(where ...interface{}) Message {
	var self Message

	//var where string
	//var binds string[]
	//for field, value := range criteria {
	//	where += field + "= ?"
	//	binds[] = value
	//}

	getDB().First(&self, where...)
	return self
}

func (this *Message) create()  {
	getDB().Create(this)
}

func (this *Message) update()  {
	getDB().Update(this)
	//getDB().Model(&this).Update("Price", 2000)
}

func (this *Message) delete()  {
	getDB().Delete(&this)
}
