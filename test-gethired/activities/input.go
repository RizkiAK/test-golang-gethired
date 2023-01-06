package activities

type InputTodolist struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type GetTodolistDetailInput struct {
	ID int `uri:"id" json:"id"`
}
