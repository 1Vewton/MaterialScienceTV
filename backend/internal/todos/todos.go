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

// NewTodo creates new todo
func NewTodo(
	todoId string,
	userId string,
	text string,
) *Todo {
	return &Todo{
		Done:   false,
		UserID: userId,
		TodoID: todoId,
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
	todoId string,
) (bool, error) {
	todos, err := gorm.G[*Todo](db).Where(
		&Todo{
			TodoID: todoId,
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
	todoId string,
) (*Todo, error) {
	HasTodo(
		ctx,
		db,
		todoId,
	)
	todo, err := gorm.G[*Todo](db).Where(
		&Todo{
			TodoID: todoId,
		},
	).First(ctx)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, fmt.Errorf(
			"todo with id %s does not exists",
			todoId,
		)
	}
	return nil, nil
}
