package entities

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	FamilyName string `json:"familyName"`
	BirthDate  string `json:"birthDate"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
