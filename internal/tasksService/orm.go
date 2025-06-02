package tasksService

type RequestBodyTask struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}

type ResponseBodyTask struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
