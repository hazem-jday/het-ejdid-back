package entities

import "time"

//Modèle du Post
type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"description"`
	Publisher string    `json:"username"`
	Image     string    `json:"image"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"`
	Source    string    `json:"source"`
}

type Like struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Article uint `json:"article"`
	User    uint `json:"user"`
}
type Dislike struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Article uint `json:"article"`
	User    uint `json:"user"`
}
