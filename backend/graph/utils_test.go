package graph

import (
	"testing"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/todos"
)

// TestToModel tests converting todos to models.Todo
func TestToModel(t *testing.T) {
	t.Parallel()
	newTodo := &todos.Todo{
		TodoID:  "114514",
		Text:    "114514",
		Done:    false,
		DueDate: time.Now(),
	}
	convertedTodo := ToSingleModel(newTodo)
	if convertedTodo.Text != newTodo.Text {
		t.Fatalf(
			"expected %s, got %s",
			newTodo.Text,
			convertedTodo.Text,
		)
	}
	if convertedTodo.Done != newTodo.Done {
		t.Fatalf(
			"expected %v, got %v",
			newTodo.Done,
			convertedTodo.Done,
		)
	}
	if convertedTodo.ID != newTodo.TodoID {
		t.Fatalf(
			"expected %s, got %s",
			newTodo.TodoID,
			convertedTodo.ID,
		)
	}
	allTodos := []*todos.Todo{
		{
			TodoID:  "114514",
			Text:    "114514",
			Done:    false,
			DueDate: time.Now(),
		},
		{
			TodoID:  "1333",
			Text:    "2245",
			Done:    false,
			DueDate: time.Now(),
		},
	}
	allModels := ToModel(
		allTodos,
	)
	if len(allModels) != 2 {
		t.Fatalf(
			"expected length of converted models to be %d, got %d",
			2,
			len(allModels),
		)
	}
}
