package graph

import (
	"context"

	"github.com/1Vewton/MaterialScienceTV/backend/graph/model"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/cookie"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/todos"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/user"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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

// VerifyUserFromCookie verifies user form cookie
func VerifyUserFromCookie(
	ctx context.Context,
	client *redis.Client,
	db *gorm.DB,
) (*user.User, bool, error) {
	request, err := cookie.GetRequestFromContext(
		ctx,
	)
	if err != nil {
		return nil, false, err
	}
	requestCookie, err := request.Cookie(
		cookie.LoginTokenKey,
	)
	if err != nil {
		return nil, false, err
	}
	token := requestCookie.Value
	return user.VerifyLoginToken(
		ctx,
		client,
		db,
		token,
	)
}
