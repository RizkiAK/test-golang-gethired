package todos

type Service interface {
	GetAll() ([]Todos, error)
	GetOne(id int) (Todos, error)
	Created(input InputTodos) (Todos, error)
	Updated(inputID int, inputData InputTodosUpdate) (Todos, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Todos, error) {
	todolist, err := s.repository.GetAll()
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) GetOne(id int) (Todos, error) {
	todolist, err := s.repository.GetOne(id)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Created(input InputTodos) (Todos, error) {
	todolist := Todos{
		Title:           input.Title,
		ActivityGroupID: input.ActivityGroupID,
		IsActive:        input.IsActive,
	}

	todolist, err := s.repository.Created(todolist)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Updated(inputID int, inputData InputTodosUpdate) (Todos, error) {
	todolist, err := s.repository.GetOne(inputID)
	if err != nil {
		return todolist, err
	}

	todolist.Title = inputData.Title
	todolist.Priority = inputData.Priority
	todolist.IsActive = inputData.IsActive

	updatedTodolist, err := s.repository.Updated(todolist)
	if err != nil {
		return updatedTodolist, err
	}

	return updatedTodolist, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
