package service

import (
	"anaia-backend/encryption"
	"anaia-backend/internal/entity"
	"anaia-backend/internal/repository"
	context "context"
	"os"
	"testing"

	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "test@exist.com", Password: encryptedPassword}

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exist.com").Return(u, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Lastname      string
		Email         string
		Password      string
		RoleId        int64
		ExpectedError error
	}{
		// Test cases. The first one is a success case, the second one is a failure case.
		{
			Name:          "BraianUser_Success",
			Lastname:      "Rosales",
			Email:         "test@test.com",
			Password:      "validPassword",
			RoleId:        1,
			ExpectedError: nil,
		},
		{
			Name:          "DaianaUse_AlreadyExists",
			Lastname:      "Caminero",
			Email:         "test@exist.com",
			Password:      "validPassword",
			RoleId:        1,
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		// Run the test cases in parallel. tc.Name is the name of the test case.
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Name, tc.Lastname, tc.Email, tc.Password, tc.RoleId)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "UserBraian_Success",
			Email:         "test@exist.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "UserDaiana_InvalidPassword",
			Email:         "test@exist.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidPassword,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
