package userService

type Users struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBodyTask struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Task    string `json:"task"`
	User_id int    `gorm:"primaryKey" json:"user_id"`
}
