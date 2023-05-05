package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:varchar(300)" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open(mysql.Open(getEnv("DBUSER") + ":" + getEnv("DBPASS") + "@tcp(localhost:" + getEnv("DBPORT") + ")/" + getEnv("DBNAME")))
	database, err := gorm.Open(mysql.Open("root:angger123@tcp(localhost:3306)/tubesStima"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
