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

var errPatient = errors.New("some error")

func TestEditPatient(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		patientRepository *mocks.MockPatientRepository
	}

	type args struct {
		fio         string
		phoneNumber string
		email       string
		insurance   string
		ID          int64
	}

	tests := []struct {
		prepare       func(f *fields)
		expectedError error
		name          string
		args          args
		wantErr       bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().EditPatient(
					int64(1),
					"test fio",
					"test number",
					"test email",
					"test insurance",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				ID:          int64(1),
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
				insurance:   "test insurance",
			},
			expectedError: nil,
			wantErr:       false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().EditPatient(
					int64(1),
					"test fio",
					"test number",
					"test email",
					"test insurance",
				).Return(int64(0), errPatient).Times(1)
			},
			args: args{
				ID:          int64(1),
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
				insurance:   "test insurance",
			},
			expectedError: errPatient,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewPatientService(
				f.patientRepository,
			)
			_, err := s.EditPatient(
				tt.args.ID,
				tt.args.fio,
				tt.args.phoneNumber,
				tt.args.email,
				tt.args.insurance,
			)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddPatient(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		patientRepository *mocks.MockPatientRepository
	}
	type args struct {
		fio         string
		phoneNumber string
		email       string
		insurance   string
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
				f.patientRepository.EXPECT().AddPatient(
					"test fio",
					"test number",
					"test email",
					"test insurance",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
				insurance:   "test insurance",
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().AddPatient(
					"test fio",
					"test number",
					"test email",
					"test insurance",
				).Return(int64(0), errPatient).Times(1)
			},
			args: args{
				fio:         "test fio",
				phoneNumber: "test number",
				email:       "test email",
				insurance:   "test insurance",
			},
			expectedError:  errPatient,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewPatientService(
				f.patientRepository,
			)
			id, err := s.AddPatient(
				tt.args.fio,
				tt.args.phoneNumber,
				tt.args.email,
				tt.args.insurance,
			)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestGetAllPatients(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		patientRepository *mocks.MockPatientRepository
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult []models.Patient
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().GetAllPatients().Return([]models.Patient{
					{
						Fio:         "test fio",
						PhoneNumber: "test phone",
						Email:       "test email",
						Insurance:   "test insurance",
						ID:          int64(1),
					},
					{
						Fio:         "test fio2",
						PhoneNumber: "test phone2",
						Email:       "test email2",
						Insurance:   "test insurance2",
						ID:          int64(2),
					},
				}, nil).Times(1)
			},
			expectedError: nil,
			expectedResult: []models.Patient{
				{
					Fio:         "test fio",
					PhoneNumber: "test phone",
					Email:       "test email",
					Insurance:   "test insurance",
					ID:          int64(1),
				},
				{
					Fio:         "test fio2",
					PhoneNumber: "test phone2",
					Email:       "test email2",
					Insurance:   "test insurance2",
					ID:          int64(2),
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().GetAllPatients().Return(nil, errPatient).Times(1)
			},
			expectedError:  errPatient,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewPatientService(
				f.patientRepository,
			)
			patients, err := s.GetAllPatients()

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, patients)
			}
		})
	}
}

func TestGetPatientByID(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		patientRepository *mocks.MockPatientRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult models.Patient
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().GetPatientByID(int64(1)).Return(models.Patient{
					Fio:         "test fio",
					PhoneNumber: "test phone",
					Email:       "test email",
					Insurance:   "test insurance",
					ID:          int64(1),
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: models.Patient{
				Fio:         "test fio",
				PhoneNumber: "test phone",
				Email:       "test email",
				Insurance:   "test insurance",
				ID:          int64(1),
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.patientRepository.EXPECT().
					GetPatientByID(int64(1)).
					Return(models.Patient{}, errPatient).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errPatient,
			expectedResult: models.Patient{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				patientRepository: mocks.NewMockPatientRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewPatientService(
				f.patientRepository,
			)
			patients, err := s.GetPatientByID(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, patients)
			}
		})
	}
}
