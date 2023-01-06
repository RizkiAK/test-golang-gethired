package todos

type Service interface {
	GetAll(activityGroupID int) ([]Todos, error)
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

func (s *service) GetAll(activityGroupID int) ([]Todos, error) {
	todolist, err := s.repository.GetAll(activityGroupID)
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
		Priority:        "very-high",
		IsActive:        "1",
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

	if inputData.Title != "" {
		todolist.Title = inputData.Title
	} else {
		if !inputData.IsActive {
			todolist.IsActive = "0"
		} else {
			todolist.IsActive = "1"
		}
	}

	todolist.Priority = "very-high"

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
