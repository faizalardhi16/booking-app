package todo

import "time"

type FormatTodos struct {
	ID        string    `json:"id"`
	Name      string    `json:"todosName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatCreateTodos(todo Todo) FormatTodos {
	t := FormatTodos{}

	t.ID = todo.ID
	t.Name = todo.Name
	t.CreatedAt = todo.CreatedAt
	t.UpdatedAt = todo.UpdatedAt

	return t
}

func FormatGetAllTodos(todo []Todo) []FormatTodos {
	t := []FormatTodos{}
	temp := FormatTodos{}

	for _, d := range todo {
		temp.ID = d.ID
		temp.Name = d.Name
		temp.CreatedAt = d.CreatedAt
		temp.UpdatedAt = d.UpdatedAt

		t = append(t, temp)
	}

	return t
}
