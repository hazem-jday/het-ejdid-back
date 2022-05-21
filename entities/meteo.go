package entities

type Meteo struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Day   string `json:"day"`
	Date  string `json:"date"`
	Image string `json:"image"`
	Min   string `json:"min"`
	Max   string `json:"max"`
}
