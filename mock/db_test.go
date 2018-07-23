package mysql

import (
	"testing"

	"bitbucket.org/xeoncross/godiapp"
)

func TestDatabase(t *testing.T) {

	db := &mockDB{}
	db.AddUser(&godiapp.User{ID: 34, Email: "john@example.com"})

	users, err := db.GetUsers()
	if err != nil {
		t.Error(err)
	}

	if len(users) != 1 && users[0].ID != 34 {
		t.Error("Failed to mock database")
	}

}
