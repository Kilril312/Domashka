package tasksService

type RequestBodyTask struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Task    string `json:"task"`
	User_id int    `json:"user_id"`
}
