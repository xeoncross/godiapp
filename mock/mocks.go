package mysql

import (
	"bitbucket.org/xeoncross/godiapp"
)

type mockDB struct {
	users []*godiapp.User
}

func (m *mockDB) AddUser(u *godiapp.User) error {
	m.users = append(m.users, u)
	return nil
}

func (m *mockDB) GetUsers() ([]*godiapp.User, error) {
	return m.users, nil
}
