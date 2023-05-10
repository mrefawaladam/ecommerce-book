package seeders

import (
	"fmt"
	"math/rand"
	"time"

	"ebook/internal/entity"

	"github.com/bxcodec/faker/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	username := viper.Get("DB_USERNAME")
	password := viper.Get("DB_PASSWORD")
	host := viper.Get("DB_HOST")
	port := viper.Get("DB_PORT")
	name := viper.Get("DB_NAME")

	rand.Seed(time.Now().UnixNano())

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migration
	err = db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Book{})
	if err != nil {
		panic("failed to auto migrate tables")
	}

	// Seed users
	users := make([]entity.User, 10)
	for i := 0; i < 10; i++ {
		err = faker.FakeData(&users[i])
		if err != nil {
			panic("failed to generate user data")
		}
		users[i].Role = "customer"
		users[i].Password = "password"

		err = db.Create(&users[i]).Error
		if err != nil {
			panic("failed to seed user data")
		}
	}

	// Seed categories
	categories := make([]entity.Category, 5)
	for i := 0; i < 5; i++ {
		err = faker.FakeData(&categories[i])
		if err != nil {
			panic("failed to generate category data")
		}

		err = db.Create(&categories[i]).Error
		if err != nil {
			panic("failed to seed category data")
		}
	}

	// Seed books
	books := make([]entity.Book, 15)
	for i := 0; i < 15; i++ {
		err = faker.FakeData(&books[i])
		if err != nil {
			panic("failed to generate book data")
		}
		books[i].CategoryID = categories[rand.Intn(len(categories))].ID
		books[i].SellerId = users[rand.Intn(len(users))].ID

		err = db.Create(&books[i]).Error
		if err != nil {
			panic("failed to seed book data")
		}
	}

	fmt.Println("Seeding completed.")
}
