package services_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/hamillka/ppo/backend/internal/services"
	"github.com/hamillka/ppo/backend/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errTimetable = errors.New("some error")

func TestGetLocationsByDoctor(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		timetableRepository *mocks.MockTimetableRepository
	}

	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		name           string
		prepare        func(f *fields)
		expectedResult []models.Office
		wantErr        bool
		args           args
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.timetableRepository.EXPECT().GetLocationsByDoctor(
					int64(1),
				).Return([]models.Office{
					{
						ID:       int64(1),
						Number:   int64(1),
						Floor:    int64(1),
						BranchID: int64(1),
					},
					{
						ID:       int64(2),
						Number:   int64(2),
						Floor:    int64(2),
						BranchID: int64(2),
					},
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: []models.Office{
				{
					ID:       int64(1),
					Number:   int64(1),
					Floor:    int64(1),
					BranchID: int64(1),
				},
				{
					ID:       int64(2),
					Number:   int64(2),
					Floor:    int64(2),
					BranchID: int64(2),
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.timetableRepository.EXPECT().GetLocationsByDoctor(
					int64(1),
				).Return([]models.Office{}, errTimetable).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errTimetable,
			expectedResult: []models.Office{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				timetableRepository: mocks.NewMockTimetableRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewTimetableService(
				f.timetableRepository,
			)
			offices, err := s.GetLocationsByDoctor(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, offices)
			}
		})
	}
}

func TestGetDoctorsByLocation(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		timetableRepository *mocks.MockTimetableRepository
	}

	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		name           string
		prepare        func(f *fields)
		expectedResult []models.Doctor
		wantErr        bool
		args           args
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.timetableRepository.EXPECT().GetDoctorsByLocation(
					int64(1),
				).Return([]models.Doctor{
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
			args: args{
				ID: int64(1),
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
				f.timetableRepository.EXPECT().GetDoctorsByLocation(
					int64(1),
				).Return([]models.Doctor{}, errTimetable).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errTimetable,
			expectedResult: []models.Doctor{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				timetableRepository: mocks.NewMockTimetableRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewTimetableService(
				f.timetableRepository,
			)
			offices, err := s.GetDoctorsByLocation(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, offices)
			}
		})
	}
}
