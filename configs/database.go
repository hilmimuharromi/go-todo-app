package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"todo-app/models"
)

var DB *gorm.DB

func DatabaseConnect() {
	env := Env
	p := Env.DatabasePort
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", env.DatabaseHost, env.DatabaseUser, env.DatabasePassword, env.DatabaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = db
	MigrateDatabase(db)
}

func MigrateDatabase(db *gorm.DB) {
	log.Println("running migrations")
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println("============ Database Migration is Error ============")
	} else {
		fmt.Println("============ Database Migration Completed ============")
	}
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
