package todo

import "gorm.io/gorm"

type Repository interface {
	CreateTodo(todo Todo) (Todo, error)
	GetAllTodo() ([]Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTodo() ([]Todo, error) {

	var todo []Todo

	err := r.db.Raw("select * from todos").Scan(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) CreateTodo(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}
