package todo

import "github.com/google/uuid"

type Service interface {
	SaveTodo(input TodoInput) (Todo, error)
	FindAllTodo() ([]Todo, error)
}

type service struct {
	repository Repository
}

func NewServiceTodo(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveTodo(input TodoInput) (Todo, error) {
	todo := Todo{}

	todo.ID = uuid.New().String()
	todo.Name = input.Name

	newTodo, err := s.repository.CreateTodo(todo)

	if err != nil {
		return newTodo, err
	}

	return newTodo, nil

}

func (s *service) FindAllTodo() ([]Todo, error) {
	todos, err := s.repository.GetAllTodo()

	if err != nil {
		return todos, err
	}

	return todos, nil
}
