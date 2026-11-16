package todos

import (
	"context"
	"fmt"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/user"
	"gorm.io/gorm"
)

// Todo defines the todo for user
type Todo struct {
	TodoID   string `gorm:"primaryKey"`
	Text     string
	DueDate  time.Time
	Done     bool
	MasterID string
	User     *user.User `gorm:"foreignkey:MasterID;references:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

// Equals checks if one todo is the same as another
func (todo *Todo) Equals(
	another *Todo,
) bool {
	return todo.TodoID == another.TodoID &&
		todo.MasterID == another.MasterID &&
		todo.Text == another.Text &&
		todo.DueDate.Equal(another.DueDate) &&
		todo.Done == another.Done
}

// NewTodo creates new todo
func NewTodo(
	todoID string,
	userID string,
	text string,
) *Todo {
	return &Todo{
		Done:     false,
		MasterID: userID,
		TodoID:   todoID,
		Text:     text,
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
	hasUser, err := user.HasUser(
		ctx,
		db,
		todo.MasterID,
	)
	if err != nil {
		return err
	}
	if !hasUser {
		return fmt.Errorf(
			"user with id %s does not exists",
			todo.MasterID,
		)
	}
	err = gorm.G[Todo](db).Create(
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
	hasTodo, err := HasTodo(
		ctx,
		db,
		todoID,
	)
	if err != nil {
		return nil, err
	}
	if !hasTodo {
		return nil, fmt.Errorf(
			"todo with id %s does not exists",
			todoID,
		)
	}
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
	return todo, nil
}

// GetTodosByUser gets todos by user id
func GetTodosByUser(
	ctx context.Context,
	db *gorm.DB,
	userID string,
) ([]*Todo, error) {
	hasUser, err := user.HasUser(
		ctx,
		db,
		userID,
	)
	if err != nil {
		return nil, err
	}
	if !hasUser {
		return nil, fmt.Errorf(
			"user with id %s does not exists",
			userID,
		)
	}
	return gorm.G[*Todo](db).Where(
		&Todo{
			MasterID: userID,
		},
	).Find(ctx)
}

// MarkTodoAsFinished marks todo as finished
func MarkTodoAsFinished(
	ctx context.Context,
	db *gorm.DB,
	userID string,
	todoID string,
) error {
	_, err := gorm.G[*Todo](db).Where(
		&Todo{
			MasterID: userID,
			TodoID:   todoID,
		},
	).Update(
		ctx,
		"Done",
		true,
	)
	return err
}

// GetTodosForRangeOfTime gets todos for a range of time about current time
func GetTodosForRangeOfTime(
	ctx context.Context,
	db *gorm.DB,
	userID string,
	timeRangeInNanoSecond int64,
) ([]*Todo, error) {
	timeStampNow := time.Now().UnixNano()
	resultTodo := []*Todo{}
	todosFetched, err := GetTodosByUser(
		ctx,
		db,
		userID,
	)
	if err != nil {
		return nil, err
	}
	for _, todo := range todosFetched {
		dueDateTimeStamp := todo.DueDate.UnixNano()
		if dueDateTimeStamp >= timeStampNow &&
			dueDateTimeStamp <= timeStampNow+timeRangeInNanoSecond {
			resultTodo = append(resultTodo, todo)
		}
	}
	return resultTodo, nil
}

// DeleteTodo deletes single todo
func DeleteTodo(
	ctx context.Context,
	userID string,
	todoID string,
	db *gorm.DB,
) error {
	_, err := gorm.G[*Todo](db).Where(
		&Todo{
			MasterID: userID,
			TodoID:   todoID,
		},
	).Delete(ctx)
	return err
}
