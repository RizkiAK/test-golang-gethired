package activites

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Activites, error)
	GetOne(id int) (Activites, error)
	Created(todolist Activites) (Activites, error)
	Updated(todolist Activites) (Activites, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Activites, error) {
	var todolist []Activites
	err := r.db.Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) GetOne(id int) (Activites, error) {
	var todolist Activites
	err := r.db.Where("id = ?", id).Find(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Created(todolist Activites) (Activites, error) {
	err := r.db.Create(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}
func (r *repository) Updated(todolist Activites) (Activites, error) {
	err := r.db.Save(&todolist).Error
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Delete(Activites{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
