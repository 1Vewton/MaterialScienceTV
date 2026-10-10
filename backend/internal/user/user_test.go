package user

import (
	"testing"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
	"github.com/google/uuid"
)

// TestUserCRUD tests the CRUD of users
func TestUserCRUD(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	db, err := database.InitDataBase(
		databasetype.Sqlite,
		":memory:",
		&User{},
	)
	if err != nil {
		t.Error(err)
	}
	id := uuid.NewString()
	userName := "114514"
	password := "abc114514"
	email := "114514@acceed.com"
	initialUser := NewUser(
		userName,
		password,
		email,
		id,
	)
	err = AddUser(
		ctx,
		db,
		initialUser,
	)
	if err != nil {
		t.Error(err)
	}
	exists, err := HasUser(
		ctx,
		db,
		id,
	)
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Errorf(
			"%s id does not exists in database after insertion",
			id,
		)
	}
	exists, err = HasUserName(
		ctx,
		db,
		"114514",
	)
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Errorf(
			"%s id does not exists in database after insertion",
			id,
		)
	}
	fetchedUser, err := GetUser(
		ctx,
		db,
		id,
	)
	if err != nil {
		t.Error(err)
	}
	if !fetchedUser.Equals(
		initialUser,
	) {
		t.Error(
			"the fetched user is not the same as the initial user",
		)
	}
	_, success, err := Login(
		ctx,
		db,
		&userName,
		nil,
		password,
	)
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error(
			"login attempt failed, it is not expected to happen",
		)
	}
	_, success, err = Login(
		ctx,
		db,
		nil,
		&email,
		password,
	)
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error(
			"login attempt failed, it is not expected to happen",
		)
	}
	_, _, err = Login(
		ctx,
		db,
		nil,
		nil,
		password,
	)
	if err == nil {
		t.Error(
			"expected to fail",
		)
	}
	_, success, err = Login(
		ctx,
		db,
		&email,
		nil,
		password,
	)
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error(
			"it is supposed to fail when logging in",
		)
	}
	_, success, err = Login(
		ctx,
		db,
		nil,
		&userName,
		password,
	)
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error(
			"it is supposed to fail when logging in",
		)
	}
}
