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
	initialUser := NewUser(
		"114514",
		"abc114514",
		".com",
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
}
