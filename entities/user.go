package entities

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string `json:"username"`
	FirstName  string `json:"firtName"`
	FamilyName string `json:"familyName"`
	BirthDate  string `json:"birthDate"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Save struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	User    uint `json:"user"`
	Article uint `json:"article"`
}
