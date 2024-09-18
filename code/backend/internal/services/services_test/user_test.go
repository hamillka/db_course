package services_test

import (
	"errors"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/hamillka/ppo/backend/internal/services"
	"github.com/hamillka/ppo/backend/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errUser = errors.New("some error")

func TestCheckUserRole(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		userRepository    *mocks.MockUserRepository
		doctorRepository  *mocks.MockDoctorRepository
		patientRepository *mocks.MockPatientRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedResult *int64
		expectedError  error
		prepare        func(f *fields)
		name           string
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.userRepository.EXPECT().CheckUserRole(
					int64(1),
				).Return(int64(0), nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(0),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.userRepository.EXPECT().CheckUserRole(
					int64(1),
				).Return(int64(-1), errUser).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errUser,
			expectedResult: pointer.ToInt64(-1),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				userRepository:    mocks.NewMockUserRepository(ctrl),
				doctorRepository:  mocks.NewMockDoctorRepository(ctrl),
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewUserService(
				f.userRepository,
				f.doctorRepository,
				f.patientRepository,
			)
			id, err := s.CheckUserRole(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		userRepository    *mocks.MockUserRepository
		doctorRepository  *mocks.MockDoctorRepository
		patientRepository *mocks.MockPatientRepository
	}
	type args struct {
		login    string
		password string
	}

	tests := []struct { //nolint: govet // impossible to fix
		expectedError  error
		prepare        func(f *fields)
		name           string
		args           args
		wantErr        bool
		expectedResult models.User
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.userRepository.EXPECT().GetUserByLoginAndPassword(
					"test login",
					"test password",
				).Return(models.User{
					Login:     "test login",
					Password:  "test password",
					PatientID: nil,
					DoctorID:  nil,
					Role:      int64(1),
					ID:        int64(1),
				}, nil).Times(1)
			},
			args: args{
				login:    "test login",
				password: "test password",
			},
			expectedError: nil,
			expectedResult: models.User{
				Login:     "test login",
				Password:  "test password",
				PatientID: nil,
				DoctorID:  nil,
				Role:      int64(1),
				ID:        int64(1),
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.userRepository.EXPECT().GetUserByLoginAndPassword(
					"test login",
					"test password",
				).Return(models.User{}, errUser).Times(1)
			},
			args: args{
				login:    "test login",
				password: "test password",
			},
			expectedError:  errUser,
			expectedResult: models.User{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				userRepository:    mocks.NewMockUserRepository(ctrl),
				doctorRepository:  mocks.NewMockDoctorRepository(ctrl),
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewUserService(
				f.userRepository,
				f.doctorRepository,
				f.patientRepository,
			)
			id, err := s.GetUserByLoginAndPassword(tt.args.login, tt.args.password)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, id)
			}
		})
	}
}
