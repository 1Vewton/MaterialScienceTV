package graph

import (
	"github.com/1Vewton/MaterialScienceTV/backend/graph/model"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/todos"
)

// ToSingleModel converts single todo to model
func ToSingleModel(
	todo *todos.Todo,
) *model.Todo {
	return &model.Todo{
		ID:   todo.TodoID,
		Text: todo.Text,
		Done: todo.Done,
		DueDateUnixTimestamp: int32(
			todo.DueDate.Unix(),
		),
	}
}

// ToModel tests the converting from todo to model.Todo
func ToModel(
	originalTodos []*todos.Todo,
) []*model.Todo {
	result := []*model.Todo{}
	for _, originalTodo := range originalTodos {
		result = append(
			result,
			ToSingleModel(
				originalTodo,
			),
		)
	}
	return result
}
