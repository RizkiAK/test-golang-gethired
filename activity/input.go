package activity

type InputTodolist struct {
	Title string
	Email string
}

type GetTodolistDetailInput struct {
	ID int `uri:"id"`
}
