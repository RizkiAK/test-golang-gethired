package activity

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Todolist, error)
	GetOne(id int) (Todolist, error)
	Created(todolist Todolist) (Todolist, error)
	Updated(todolist Todolist) (Todolist, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Todolist, error) {
	var todolist []Todolist
	err := r.db.Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) GetOne(id int) (Todolist, error) {
	var todolist Todolist
	err := r.db.Where("id = ?", id).Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Created(todolist Todolist) (Todolist, error) {
	err := r.db.Create(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}
func (r *repository) Updated(todolist Todolist) (Todolist, error) {
	err := r.db.Save(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Delete(Todolist{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
