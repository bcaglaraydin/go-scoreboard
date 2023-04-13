package database

import (
	"fmt"
)

// type DbInstance struct {
// 	Db *gorm.DB
// }

// var DB DbInstance

func ConnectDb() {

	// dsn := fmt.Sprintf(
	// 	"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Istanbul",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	// helpers.LogError("Failed to connect to dataabase!", err)
	// log.Println("Connected to database")

	// db.Logger = logger.Default.LogMode(logger.Info)
	// log.Println("Running migrations..")
	// db.AutoMigrate(&models.User{})

	// DB = DbInstance{
	// 	Db: db,
	// }
	fmt.Printf("connected")
}
