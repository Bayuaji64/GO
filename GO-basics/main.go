package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	// Membuat DSN dari nilai .env
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s", dbUser, dbPassword, dbHost, dbName, dbCharset, dbParseTime, dbLoc)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect with the database")
	}

	// db.Migrator().DropTable(&User{}, &Address{}, &Book{})
	// db.Migrator().CreateTable(&User{}, &Address{}, &Book{})
	db.AutoMigrate(&User{}, &Book{}, &Address{})

	// userCreate := User{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Email:     "john@doe.com",
	// }

	// db.Create(&userCreate)

	// userUpdate := User{
	// 	Id:        1,
	// 	FirstName: "John 2",
	// 	LastName:  "Doe 2",
	// 	Email:     "john@doe2.com",
	// }

	// db.Updates(&userUpdate)

	// userDelete := User{
	// 	Id: 1,
	// }

	// db.Delete(&userDelete)

	// users := []User{
	// 	{
	// 		FirstName: "John",
	// 		LastName:  "Doe",
	// 		Email:     "John@mail.com",
	// 	},
	// 	{
	// 		FirstName: "John1",
	// 		LastName:  "Doe1",
	// 		Email:     "John1@mail.com",
	// 	},
	// 	{
	// 		FirstName: "John2",
	// 		LastName:  "Doe2",
	// 		Email:     "John2@mail.com",
	// 	},
	// }

	// for _, v := range users {

	// 	db.Create(&v)

	// }

	//cek data

	// user := User{}

	// db.Last(&user)
	// db.Where("last_name", "Doe1").First(&user)

	// fmt.Println(user)

	// user := User{
	// 	Model: gorm.Model{
	// 		CreatedAt: time.Now(),
	// 	},
	// }

	// fmt.Println(user)

	// user := User1{
	// 	Email: sql.NullString{
	// 		String: "ada@.com",
	// 		Valid:  true,
	// 	},
	// }

	// db.Create(&user)

	user := User{
		FirstName: "ada",
		LastName:  "adatidak",
		Email:     "ada@.com",
		Address: Address{
			Name: "Main str.",
		},
		Books: []Book{
			{
				Title: "INI BUKU",
			},
		},
	}

	db.Create(&user)

	// u := User{}

	// fmt.Println(u)

}

type User1 struct {
	gorm.Model
	FirstName sql.NullString `gorm:"type:VARCHAR(30); null"`
	LastName  sql.NullString `gorm:"size:100; default:'Smith'"`
	Email     sql.NullString `gorm:"unique; not null"`
}
type User struct {
	gorm.Model
	FirstName string  `gorm:"type:VARCHAR(30); null"`
	LastName  string  `gorm:"size:100; default:'Smith'"`
	Email     string  `gorm:"unique; not null"`
	Address   Address `gorm:"foreignkey:UserId"`
	Books     []Book  `gorm:"many2many:user_books"`
}

type Address struct {
	gorm.Model
	UserId uint
	Name   string
}
type Book struct {
	ID    uint
	Title string
}
