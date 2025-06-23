package userService

type Users struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
