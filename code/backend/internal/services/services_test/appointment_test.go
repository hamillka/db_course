package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/hamillka/ppo/backend/internal/services"
	"github.com/hamillka/ppo/backend/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errAppointment = errors.New("some error")

func TestCreateAppointment(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		appointmentRepository *mocks.MockAppointmentRepository
	}
	type args struct {
		dateTime  time.Time
		ID        int64
		patientID int64
		doctorID  int64
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
				f.appointmentRepository.EXPECT().CreateAppointment(
					int64(1),
					int64(1),
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				ID:        int64(1),
				patientID: int64(1),
				doctorID:  int64(1),
				dateTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().CreateAppointment(
					int64(1),
					int64(1),
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				).Return(int64(0), errAppointment).Times(1)
			},
			args: args{
				ID:        int64(1),
				patientID: int64(1),
				doctorID:  int64(1),
				dateTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError:  errAppointment,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				appointmentRepository: mocks.NewMockAppointmentRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewAppointmentService(
				f.appointmentRepository,
			)
			id, err := s.CreateAppointment(tt.args.patientID, tt.args.doctorID, tt.args.dateTime)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestCancelAppointment(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		appointmentRepository *mocks.MockAppointmentRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedResult *int64
		prepare        func(f *fields)
		expectedError  error
		name           string
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().CancelAppointment(
					int64(1),
				).Return(nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().CancelAppointment(
					int64(1),
				).Return(errAppointment).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errAppointment,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				appointmentRepository: mocks.NewMockAppointmentRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewAppointmentService(
				f.appointmentRepository,
			)
			err := s.CancelAppointment(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetAppointmentsByPatient(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		appointmentRepository *mocks.MockAppointmentRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		name           string
		expectedError  error
		prepare        func(f *fields)
		expectedResult []*models.Appointment
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().GetAppointmentsByPatient(
					int64(1),
				).Return([]*models.Appointment{
					{
						DateTime:  time.Date(2024, 12, 18, 16, 45, 0, 0, time.UTC),
						ID:        int64(1),
						PatientID: int64(1),
						DoctorID:  int64(1),
					},
					{
						DateTime:  time.Date(2024, 12, 25, 11, 15, 0, 0, time.UTC),
						ID:        int64(2),
						PatientID: int64(1),
						DoctorID:  int64(2),
					},
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: []*models.Appointment{
				{
					DateTime:  time.Date(2024, 12, 18, 16, 45, 0, 0, time.UTC),
					ID:        int64(1),
					PatientID: int64(1),
					DoctorID:  int64(1),
				},
				{
					DateTime:  time.Date(2024, 12, 25, 11, 15, 0, 0, time.UTC),
					ID:        int64(2),
					PatientID: int64(1),
					DoctorID:  int64(2),
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().GetAppointmentsByPatient(
					int64(1),
				).Return(nil, errAppointment).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errAppointment,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				appointmentRepository: mocks.NewMockAppointmentRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewAppointmentService(
				f.appointmentRepository,
			)
			appointments, err := s.GetAppointmentsByPatient(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, appointments)
			}
		})
	}
}

func TestGetAppointmentsByDoctor(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		appointmentRepository *mocks.MockAppointmentRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		name           string
		expectedError  error
		prepare        func(f *fields)
		expectedResult []*models.Appointment
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().GetAppointmentsByDoctor(
					int64(1),
				).Return([]*models.Appointment{
					{
						DateTime:  time.Date(2024, 11, 18, 16, 45, 0, 0, time.UTC),
						ID:        int64(1),
						PatientID: int64(4),
						DoctorID:  int64(1),
					},
					{
						DateTime:  time.Date(2024, 11, 25, 11, 15, 0, 0, time.UTC),
						ID:        int64(2),
						PatientID: int64(5),
						DoctorID:  int64(1),
					},
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: []*models.Appointment{
				{
					DateTime:  time.Date(2024, 11, 18, 16, 45, 0, 0, time.UTC),
					ID:        int64(1),
					PatientID: int64(4),
					DoctorID:  int64(1),
				},
				{
					DateTime:  time.Date(2024, 11, 25, 11, 15, 0, 0, time.UTC),
					ID:        int64(2),
					PatientID: int64(5),
					DoctorID:  int64(1),
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.appointmentRepository.EXPECT().GetAppointmentsByDoctor(
					int64(1),
				).Return(nil, errAppointment).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errAppointment,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				appointmentRepository: mocks.NewMockAppointmentRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewAppointmentService(
				f.appointmentRepository,
			)
			appointments, err := s.GetAppointmentsByDoctor(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, appointments)
			}
		})
	}
}
