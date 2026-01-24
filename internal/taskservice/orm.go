package taskservice

type TaskaRequest struct {
	Task string `json:"task"`
}

type Taska struct {
	Task string `json:"task"`
	ID   string `gorm:"primarykey" json:"id"`
}
