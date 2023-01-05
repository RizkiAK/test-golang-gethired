package activites

type Service interface {
	GetAll() ([]Activites, error)
	GetOne(id int) (Activites, error)
	Created(input InputTodolist) (Activites, error)
	Updated(inputID int, inputData InputTodolist) (Activites, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Activites, error) {
	todolist, err := s.repository.GetAll()
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) GetOne(id int) (Activites, error) {
	todolist, err := s.repository.GetOne(id)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Created(input InputTodolist) (Activites, error) {
	todolist := Activites{
		Title: input.Title,
		Email: input.Email,
	}

	todolist, err := s.repository.Created(todolist)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (s *service) Updated(inputID int, inputData InputTodolist) (Activites, error) {
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
