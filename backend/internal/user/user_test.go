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
	err = AddUser(
		ctx,
		db,
		&User{
			UserID:   id,
			UserName: "Tester",
			Password: "1145141919810",
		},
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
}
