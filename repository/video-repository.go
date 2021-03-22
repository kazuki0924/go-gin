package repository

import (
	"log"

	"github.com/kazuki0924/go-gin/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{connection: db}
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		log.Fatalln("failed to connect to database")
	}
	defer sqlDB.Close()
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Joins("Author").Find(&videos)
	return videos
}
