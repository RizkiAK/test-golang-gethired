package activity

type Service interface {
	GetAll() ([]Todolist, error)
	GetOne(id int) (Todolist, error)
	Created(input InputTodolist) (Todolist, error)
	Updated(inputID int, inputData InputTodolist) (Todolist, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Todolist, error) {
	todolist, err := s.repository.GetAll()
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) GetOne(id int) (Todolist, error) {
	todolist, err := s.repository.GetOne(id)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Created(input InputTodolist) (Todolist, error) {
	todolist := Todolist{
		Title: input.Title,
		Email: input.Email,
	}

	todolist, err := s.repository.Created(todolist)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Updated(inputID int, inputData InputTodolist) (Todolist, error) {
	todolist, err := s.repository.GetOne(inputID)
	if err != nil {
		return todolist, err
	}

	todolist.Title = inputData.Title

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
