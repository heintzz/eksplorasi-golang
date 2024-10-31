// FLOW: implement repository tadi nah kalo di-mockinig you tinggal give data aja
// ga usah ngambil dari db, tp tetep pastiin sama yah
package services

import (
	"database/sql"
	"testing"
)

// ga bener2 ngehit db

type mockRepo struct{}

func newMockRepo() mockRepo {
	return mockRepo{}
}

func (m mockRepo) GetAll() ([]User, error) {
	return getAll()
}

var (
	getAll func() ([]User, error)
)

var svc Service

func init() {
	repo := newMockRepo()
	svc = NewService(repo)
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {		
		getAll = func() ([]User, error) {
			return []User{
				{
					Id:   1,
					Name: "NooBeeID",
				},
			}, nil
		}

		users, err := svc.GetAll()
		if err != nil {
			t.Errorf("expect no error, but got : %v", err.Error())
		}

		if len(users) == 0 {
			t.Error("expect len users > 0")
		}
	})

	t.Run("fail_notfound", func(t *testing.T) {
		getAll = func() ([]User, error) {
			return nil, sql.ErrNoRows
		}

		users, err := svc.GetAll()
		if err == nil {
			t.Error("expect error not found, but got nil")
		}

		if len(users) != 0 {
			t.Error("expect len users == 0")
		}
	})
}
