package activities

type Service interface {
	GetAll() ([]Activities, error)
	GetOne(id int) (Activities, error)
	Created(input InputTodolist) (Activities, error)
	Updated(inputID int, inputData InputTodolist) (Activities, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Activities, error) {
	todolist, err := s.repository.GetAll()
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) GetOne(id int) (Activities, error) {
	todolist, err := s.repository.GetOne(id)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Created(input InputTodolist) (Activities, error) {
	todolist := Activities{
		Title: input.Title,
		Email: input.Email,
	}

	todolist, err := s.repository.Created(todolist)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Updated(inputID int, inputData InputTodolist) (Activities, error) {
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
