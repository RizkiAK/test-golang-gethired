package activities

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Activities, error)
	GetOne(id int) (Activities, error)
	Created(todolist Activities) (Activities, error)
	Updated(todolist Activities) (Activities, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Activities, error) {
	var todolist []Activities
	err := r.db.Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) GetOne(id int) (Activities, error) {
	var todolist Activities
	err := r.db.Where("id = ?", id).Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Created(todolist Activities) (Activities, error) {
	err := r.db.Create(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}
func (r *repository) Updated(todolist Activities) (Activities, error) {
	err := r.db.Save(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Delete(Activities{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
