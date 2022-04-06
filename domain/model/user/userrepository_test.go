package user

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func Test_FindByUserName(t *testing.T) {
	userName, _ := NewUserName("userName")
	userId, _ := NewUserId("userId")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM users WHERE name = $1`)).
		WithArgs("userName").
		WillReturnRows(mock.NewRows([]string{"userId", "userName"}).AddRow("userId", "userName"))
	mock.ExpectCommit()

	got, err := userRepository.FindByUserName(userName)
	if err != nil {
		t.Error(err)
	}
	want := &User{id: *userId, name: *userName}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(User{}, UserName{}, UserId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
