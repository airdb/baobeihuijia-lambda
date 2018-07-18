package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Article struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID      string
	AvatarUrl string
	Nickname  string
	Gender    uint
	Province  string
	City      string
	Country   string
	Address   string
	// Title      string `gorm:"size:60"`
	Title      string
	BirthedAt  time.Time
	MissedAt   time.Time
	Arcid      string // ¹浵±à
	Age        int
	Characters string
	Details    string
}

func AddArticle(article Article) (UUID string) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Article{})

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	article.MissedAt = time.Now()

	UUID = uuid.New().String()
	article.UUID = UUID
	// fmt.Println("db=====", article)
	conn.Create(&article)
	// db.Save(&lbi)
	return
}

func GetArticle(uuid string) (article Article, err error) {

	fmt.Println(uuid)
	conn.Where("uuid = ?", uuid).Last(&article)

	fmt.Println("%s", article)
	return
}

func GetArticles() (articles []Article) {
	// var lbis []LostBabyInfo
	conn.Limit(5).Find(&articles)

	fmt.Println("%s", articles)
	return
}

func UpdateArticle() {
}
func DeleteArticle(uuid string) (flag bool) {
	// var lbis []LostBabyInfo
	conn.Where("uuid = ?", uuid).Delete(Article{})

	return
}
