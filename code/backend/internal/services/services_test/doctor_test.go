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

var errDoctor = errors.New("some error")

func TestEditDoctor(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		doctorRepository *mocks.MockDoctorRepository
	}

	type args struct {
		fio         string
		phoneNumber string
		email       string
		ID          int64
	}

	tests := []struct {
		prepare        func(f *fields)
		expectedResult *int64
		expectedError  error
		name           string
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().EditDoctor(
					int64(1),
					"test fio",
					"test number",
					"test email",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				ID:          int64(1),
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().EditDoctor(
					int64(1),
					"test fio",
					"test number",
					"test email",
				).Return(int64(0), errDoctor).Times(1)
			},
			args: args{
				ID:          int64(1),
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
			},
			expectedError:  errDoctor,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				doctorRepository: mocks.NewMockDoctorRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewDoctorService(
				f.doctorRepository,
			)
			id, err := s.EditDoctor(tt.args.ID, tt.args.fio, tt.args.phoneNumber, tt.args.email)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestAddDoctor(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		doctorRepository *mocks.MockDoctorRepository
	}
	type args struct {
		fio         string
		phoneNumber string
		email       string
	}

	tests := []struct {
		expectedResult *int64
		expectedError  error
		name           string
		prepare        func(f *fields)
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().AddDoctor(
					"test fio",
					"test number",
					"test email",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().AddDoctor(
					"test fio",
					"test number",
					"test email",
				).Return(int64(0), errDoctor).Times(1)
			},
			args: args{
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
			},
			expectedError:  errDoctor,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				doctorRepository: mocks.NewMockDoctorRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewDoctorService(
				f.doctorRepository,
			)
			id, err := s.AddDoctor(tt.args.fio, tt.args.phoneNumber, tt.args.email)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestGetAllDoctors(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		doctorRepository *mocks.MockDoctorRepository
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult []models.Doctor
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().GetAllDoctors().Return([]models.Doctor{
					{
						Fio:         "test fio",
						PhoneNumber: "test phone",
						Email:       "test email",
						ID:          int64(1),
					},
					{
						Fio:         "test fio2",
						PhoneNumber: "test phone2",
						Email:       "test email2",
						ID:          int64(2),
					},
				}, nil).Times(1)
			},
			expectedError: nil,
			expectedResult: []models.Doctor{
				{
					Fio:         "test fio",
					PhoneNumber: "test phone",
					Email:       "test email",
					ID:          int64(1),
				},
				{
					Fio:         "test fio2",
					PhoneNumber: "test phone2",
					Email:       "test email2",
					ID:          int64(2),
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().GetAllDoctors().Return(nil, errDoctor).Times(1)
			},
			expectedError:  errDoctor,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				doctorRepository: mocks.NewMockDoctorRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewDoctorService(
				f.doctorRepository,
			)
			doctors, err := s.GetAllDoctors()

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, doctors)
			}
		})
	}
}

func TestGetDoctorByID(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		doctorRepository *mocks.MockDoctorRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult models.Doctor
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().GetDoctorByID(int64(1)).Return(models.Doctor{
					Fio:         "test fio",
					PhoneNumber: "test phone",
					Email:       "test email",
					ID:          int64(1),
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: models.Doctor{
				Fio:         "test fio",
				PhoneNumber: "test phone",
				Email:       "test email",
				ID:          int64(1),
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.doctorRepository.EXPECT().GetDoctorByID(int64(1)).Return(models.Doctor{}, errDoctor).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errDoctor,
			expectedResult: models.Doctor{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				doctorRepository: mocks.NewMockDoctorRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewDoctorService(
				f.doctorRepository,
			)
			doctors, err := s.GetDoctorByID(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, doctors)
			}
		})
	}
}
