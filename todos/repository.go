package todos

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Todos, error)
	GetOne(id int) (Todos, error)
	Created(todolist Todos) (Todos, error)
	Updated(todolist Todos) (Todos, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Todos, error) {
	var todolist []Todos
	err := r.db.Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) GetOne(id int) (Todos, error) {
	var todolist Todos
	err := r.db.Where("id = ?", id).Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Created(todolist Todos) (Todos, error) {
	err := r.db.Create(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}
func (r *repository) Updated(todolist Todos) (Todos, error) {
	err := r.db.Save(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Delete(Todos{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
