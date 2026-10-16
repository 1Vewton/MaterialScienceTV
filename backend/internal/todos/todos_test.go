package todos

import (
	"testing"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
	"github.com/google/uuid"
)

// TestTimeSetting tests the setting of time
func TestTimeSetting(t *testing.T) {
	t.Parallel()
	userID := uuid.NewString()
	taskID := uuid.NewString()
	newTodo := NewTodo(
		taskID,
		userID,
		"test",
	)
	testTime := time.Now()
	newTodo.SetDueDateByTime(
		testTime,
	)
	newTodo2 := NewTodo(
		taskID,
		userID,
		"test",
	)
	newTodo2.SetDueDateByTimestamp(
		testTime.UnixNano(),
	)
	if !newTodo.DueDate.Equal(newTodo2.DueDate) {
		t.Errorf(
			"newTodo: %s, newTodo2: %s",
			newTodo.DueDate.UTC().String(),
			newTodo2.DueDate.UTC().String(),
		)
	}
}

// TestTodosCRUD tests the CRUD for todos
func TestTodosCRUD(t *testing.T) {
	t.Parallel()
	// Prepare data
	ctx := t.Context()
	db, err := database.InitDataBase(
		databasetype.Sqlite,
		":memory:",
		&Todo{},
	)
	if err != nil {
		t.Error(err)
	}
	// Define experiment variables
	userID := uuid.NewString()
	taskID := uuid.NewString()
	newTodo := NewTodo(
		taskID,
		userID,
		"test",
	)
	testTime := time.Now()
	newTodo.SetDueDateByTime(
		testTime,
	)
	// CRUD
	err = SaveTodo(
		ctx,
		db,
		newTodo,
	)
	if err != nil {
		t.Error(err)
	}
	hasTodo, err := HasTodo(
		ctx,
		db,
		taskID,
	)
	if err != nil {
		t.Error(err)
	}
	if !hasTodo {
		t.Errorf(
			"todo with id %s expected to exists, however not exists",
			taskID,
		)
	}
}
