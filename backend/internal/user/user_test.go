package user

import (
	"testing"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/redismanager"
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
	client := redismanager.NewRedisConfig(
		"localhost:6379",
		"",
	).ToClient()
	if err != nil {
		t.Error(err)
	}
	// CRUD
	id := uuid.NewString()
	userName := "114514"
	password := "abc114514"
	email := "114514@acceed.com"
	initialUser, err := NewUser(
		userName,
		password,
		email,
		id,
	)
	if err != nil {
		t.Error(err)
	}
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
	// Login
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
	user, success, err := Login(
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
	if err == nil {
		t.Error(
			"expected to crush here",
		)
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
	if err == nil {
		t.Error(
			"expected to crush here",
		)
	}
	if success {
		t.Error(
			"it is supposed to fail when logging in",
		)
	}
	if user == nil {
		t.Error(
			"a nil user is returned",
		)
	}
	token := uuid.NewString()
	err = UploadLoginToken(
		ctx,
		client,
		token,
		user.UserID,
		5*time.Minute,
	)
	if err != nil {
		t.Error(err)
	}
	fetchedUser, verified, err := VerifyLoginToken(
		ctx,
		client,
		db,
		token,
	)
	if err != nil {
		t.Error(err)
	}
	if !fetchedUser.Equals(user) {
		t.Error(
			"the user data fetched is not equal to the original user",
		)
	}
	if !verified {
		t.Errorf(
			"cannot verify %s",
			token,
		)
	}
}
