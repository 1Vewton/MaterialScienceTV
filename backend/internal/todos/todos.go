package todos

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Todo defines the todo for user
type Todo struct {
	gorm.Model
	TodoID  string
	UserID  string
	Text    string
	DueDate time.Time
	Done    bool
}

// Equals checks if one todo is the same as another
func (todo *Todo) Equals(
	another *Todo,
) bool {
	return todo.TodoID == another.TodoID
}

// NewTodo creates new todo
func NewTodo(
	todoID string,
	userID string,
	text string,
) *Todo {
	return &Todo{
		Done:   false,
		UserID: userID,
		TodoID: todoID,
		Text:   text,
	}
}

// SetDueDateByTime sets due date by time.Time
func (todo *Todo) SetDueDateByTime(
	dueDate time.Time,
) *Todo {
	todo.DueDate = dueDate
	return todo
}

// SetDueDateByTimestamp sets due date by timestamp in nanosecond
func (todo *Todo) SetDueDateByTimestamp(
	timeStamp int64,
) *Todo {
	todo.DueDate = time.Unix(
		0,
		timeStamp,
	)
	return todo
}

// SaveTodo saves todo in the database
func SaveTodo(
	ctx context.Context,
	db *gorm.DB,
	todo *Todo,
) error {
	err := gorm.G[Todo](db).Create(
		ctx,
		todo,
	)
	return err
}

// HasTodo checks if certain todo exists
func HasTodo(
	ctx context.Context,
	db *gorm.DB,
	todoID string,
) (bool, error) {
	todos, err := gorm.G[*Todo](db).Where(
		&Todo{
			TodoID: todoID,
		},
	).Find(ctx)
	if err != nil {
		return false, err
	}
	if len(todos) < 1 {
		return false, nil
	}
	return true, nil
}

// GetTodo gets info from the todo database using todo id
func GetTodo(
	ctx context.Context,
	db *gorm.DB,
	todoID string,
) (*Todo, error) {
	HasTodo(
		ctx,
		db,
		todoID,
	)
	todo, err := gorm.G[*Todo](db).Where(
		&Todo{
			TodoID: todoID,
		},
	).First(ctx)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, fmt.Errorf(
			"todo with id %s does not exists",
			todoID,
		)
	}
	return nil, nil
}
