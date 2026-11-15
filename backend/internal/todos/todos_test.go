package todos

import (
	"testing"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/user"
	"github.com/google/uuid"
)

// TestTimeSetting tests the setting of time
func TestTimeSetting(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	db, err := database.InitDataBase(
		databasetype.Sqlite,
		":memory:",
		&user.User{},
		&Todo{},
	)
	if err != nil {
		t.Fatal(err)
	}
	userID := uuid.NewString()
	taskID := uuid.NewString()
	newUser, err := user.NewUser(
		"name1",
		"abc@114514",
		"1234@163.com",
		userID,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = user.AddUser(
		ctx,
		db,
		newUser,
	)
	if err != nil {
		t.Fatal(err)
	}
	newTodo, err := NewTodo(
		ctx,
		db,
		taskID,
		userID,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
	testTime := time.Now()
	newTodo.SetDueDateByTime(
		testTime,
	)
	newTodo2, err := NewTodo(
		ctx,
		db,
		taskID,
		userID,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
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
		":memory:?_pragma=foreign_keys(1)",
		&Todo{},
		&user.User{},
	)
	if err != nil {
		t.Error(err)
	}
	// Define experiment variables
	userID := uuid.NewString()
	taskID := uuid.NewString()
	newUser, err := user.NewUser(
		"name1",
		"abc@114514",
		"1234@163.com",
		userID,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = user.AddUser(
		ctx,
		db,
		newUser,
	)
	if err != nil {
		t.Fatal(err)
	}
	newTodo, err := NewTodo(
		ctx,
		db,
		taskID,
		userID,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
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

// TestForeginKey tests the connection between todos and users
func TestForeginKey(t *testing.T) {
	t.Parallel()
	// Prepare data
	ctx := t.Context()
	db, err := database.InitDataBase(
		databasetype.Sqlite,
		":memory:?_pragma=foreign_keys(1)",
		&Todo{},
		&user.User{},
	)
	if err != nil {
		t.Fatal(err)
	}
	userID := uuid.NewString()
	todoID := uuid.NewString()
	// Add user
	newUser, err := user.NewUser(
		"name1",
		"abc@114514",
		"1234@163.com",
		userID,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = user.AddUser(
		ctx,
		db,
		newUser,
	)
	if err != nil {
		t.Fatal(err)
	}
	newTodo, err := NewTodo(
		ctx,
		db,
		todoID,
		userID,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveTodo(
		ctx,
		db,
		newTodo,
	)
	if err != nil {
		t.Fatal(
			err,
		)
	}
	hasTodo, err := HasTodo(
		ctx,
		db,
		todoID,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !hasTodo {
		t.Fatalf(
			"expected to have todo with id %s",
			todoID,
		)
	}
	cnt := db.Model(newTodo).Association("User").Count()
	if cnt != 1 {
		t.Fatalf(
			"expected to connect %d outs, got %d",
			1,
			cnt,
		)
	}
	err = user.DeleteUser(
		ctx,
		db,
		userID,
	)
	if err != nil {
		t.Error(err)
	}
	hasUser, err := user.HasUser(
		ctx,
		db,
		userID,
	)
	if err != nil {
		t.Fatal(err)
	}
	if hasUser {
		t.Fatalf(
			"user with id %s still exists",
			userID,
		)
	}
	hasTodoNew, err := HasTodo(
		ctx,
		db,
		todoID,
	)
	if err != nil {
		t.Fatal(err)
	}
	if hasTodoNew {
		t.Fatalf(
			"expected todo %s to get deleted with the deletion of user %s",
			todoID,
			userID,
		)
	}
}
