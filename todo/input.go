package todo

import validation "github.com/go-ozzo/ozzo-validation"

type TodoInput struct {
	Name string `json:"name" validate:"required"`
}

func (p TodoInput) ValidateTodoInput() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
	)
}
