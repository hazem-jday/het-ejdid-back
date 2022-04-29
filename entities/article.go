package entities

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Publisher string    `json:"publisher"`
	Image     string    `json:"image"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"`
	Source    string    `json:"source"`
}

type NewsHighlights struct {
	Nat   []Article `json:"nat"`
	Inter []Article `json:"inter"`
	Sport []Article `json:"sport"`
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

type Save struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	User    uint `json:"user"`
	Article uint `json:"article"`
}
