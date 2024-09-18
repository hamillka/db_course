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

var errOffice = errors.New("some error")

func TestEditOffice(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		officeRepository *mocks.MockOfficeRepository
	}

	type args struct {
		ID       int64
		number   int64
		floor    int64
		branchID int64
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
				f.officeRepository.EXPECT().EditOffice(
					int64(1),
					int64(1),
					int64(1),
					int64(1),
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				ID:       int64(1),
				number:   int64(1),
				floor:    int64(1),
				branchID: int64(1),
			},
			expectedError: nil,
			wantErr:       false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.officeRepository.EXPECT().EditOffice(
					int64(1),
					int64(1),
					int64(1),
					int64(1),
				).Return(int64(0), errOffice).Times(1)
			},
			args: args{
				ID:       int64(1),
				number:   int64(1),
				floor:    int64(1),
				branchID: int64(1),
			},
			expectedError: errOffice,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				officeRepository: mocks.NewMockOfficeRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewOfficeService(
				f.officeRepository,
			)
			_, err := s.EditOffice(tt.args.ID, tt.args.number, tt.args.floor, tt.args.branchID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddOffice(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		officeRepository *mocks.MockOfficeRepository
	}
	type args struct {
		number   int64
		floor    int64
		branchID int64
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
				f.officeRepository.EXPECT().AddOffice(
					int64(1),
					int64(1),
					int64(1),
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				number:   int64(1),
				floor:    int64(1),
				branchID: int64(1),
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.officeRepository.EXPECT().AddOffice(
					int64(1),
					int64(1),
					int64(1),
				).Return(int64(0), errOffice).Times(1)
			},
			args: args{
				number:   int64(1),
				floor:    int64(1),
				branchID: int64(1),
			},
			expectedError:  errOffice,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				officeRepository: mocks.NewMockOfficeRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewOfficeService(
				f.officeRepository,
			)
			id, err := s.AddOffice(tt.args.number, tt.args.floor, tt.args.branchID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestGetAllOffices(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		officeRepository *mocks.MockOfficeRepository
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult []models.Office
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.officeRepository.EXPECT().GetAllOffices().Return([]models.Office{
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
				f.officeRepository.EXPECT().GetAllOffices().Return(nil, errOffice).Times(1)
			},
			expectedError:  errOffice,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				officeRepository: mocks.NewMockOfficeRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewOfficeService(
				f.officeRepository,
			)
			offices, err := s.GetAllOffices()

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, offices)
			}
		})
	}
}

func TestGetOfficeByID(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		officeRepository *mocks.MockOfficeRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult models.Office
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.officeRepository.EXPECT().GetOfficeByID(int64(1)).Return(models.Office{
					ID:       int64(1),
					Number:   int64(1),
					Floor:    int64(1),
					BranchID: int64(1),
				}, nil).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError: nil,
			expectedResult: models.Office{
				ID:       int64(1),
				Number:   int64(1),
				Floor:    int64(1),
				BranchID: int64(1),
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.officeRepository.EXPECT().GetOfficeByID(int64(1)).Return(models.Office{}, errOffice).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errOffice,
			expectedResult: models.Office{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				officeRepository: mocks.NewMockOfficeRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewOfficeService(
				f.officeRepository,
			)
			offices, err := s.GetOfficeByID(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, offices)
			}
		})
	}
}
