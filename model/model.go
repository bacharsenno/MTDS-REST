package model

import "github.com/jinzhu/gorm"

type Login struct {
	Username string
	Password string
}

type Teacher struct {
	ID          int    `gorm:"PRIMARY_KEY, AUTOINCREMENT" form:"ID" json:"ID"`
	Firstname   string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname    string `gorm:"not null" form:"lastname" json:"lastname"`
	Email       string `gorm:"not null" form:"email" json:"email"`
	Username    string `gorm:"not null" form:"username" json:"username"`
	Password    string `gorm:"not null" form:"password" json:"password"`
	PhoneNumber string `gorm:"not null" form:"phonenumber" json:"phonenumber"`
}

type Parent struct {
	ID          int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Firstname   string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname    string `gorm:"not null" form:"lastname" json:"lastname"`
	Email       string `gorm:"not null" form:"email" json:"email"`
	Username    string `gorm:"not null" form:"username" json:"username"`
	Password    string `gorm:"not null" form:"password" json:"password"`
	PhoneNumber string `gorm:"not null" form:"phonenumber" json:"phonenumber"`
}

type Subject struct {
	ID    int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Name  string `gorm:"not null" form:"name" json:"name"`
	Class int    `gorm:"not null" form:"class" json:"class"`
}

type Class struct {
	ID       int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Name     string `gorm:"not null" form:"name" json:"name"`
	Location string `gorm:"not null" form:"location" json:"location"`
}

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("mysql", "root:@/testdb")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&Teacher{}) {
		db.CreateTable(&Teacher{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Teacher{})
	}
	if !db.HasTable(&Parent{}) {
		db.CreateTable(&Parent{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Parent{})
	}
	if !db.HasTable(&Subject{}) {
		db.CreateTable(&Subject{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Subject{})
	}
	if !db.HasTable(&Class{}) {
		db.CreateTable(&Class{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Class{})
	}
	return db
}
