package tasksService

type RequestBodyTask struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}
