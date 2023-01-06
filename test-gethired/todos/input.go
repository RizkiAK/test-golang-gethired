package todos

type InputTodos struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

type GetTodosDetailInput struct {
	ID int `uri:"id" json:"id"`
}

type InputTodosUpdate struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
	Status   string `json:"status"`
}
