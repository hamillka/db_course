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

var errBranch = errors.New("some error")

func TestEditBranch(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		branchRepository *mocks.MockBranchRepository
	}
	type args struct {
		name        string
		address     string
		phoneNumber string
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
				f.branchRepository.EXPECT().EditBranch(
					int64(1),
					"test name",
					"test address",
					"test number",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				ID:          int64(1),
				name:        "test name",
				address:     "test address",
				phoneNumber: "test number",
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().EditBranch(
					int64(1),
					"test name",
					"test address",
					"test number",
				).Return(int64(0), errBranch).Times(1)
			},
			args: args{
				ID:          int64(1),
				name:        "test name",
				address:     "test address",
				phoneNumber: "test number",
			},
			expectedError: errBranch,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				branchRepository: mocks.NewMockBranchRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewBranchService(
				f.branchRepository,
			)
			id, err := s.EditBranch(tt.args.ID, tt.args.name, tt.args.address, tt.args.phoneNumber)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestAddBranch(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		branchRepository *mocks.MockBranchRepository
	}
	type args struct {
		name        string
		address     string
		phoneNumber string
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
				f.branchRepository.EXPECT().AddBranch(
					"testBranch",
					"example street, 11",
					"+7 (123)",
				).Return(int64(1), nil).Times(1)
			},
			args: args{
				name:        "testBranch",
				address:     "example street, 11",
				phoneNumber: "+7 (123)",
			},
			expectedError:  nil,
			expectedResult: pointer.ToInt64(1),
			wantErr:        false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().AddBranch(
					"testBranch",
					"example street, 11",
					"+7 (123)",
				).Return(int64(0), errBranch).Times(1)
			},
			args: args{
				name:        "testBranch",
				address:     "example street, 11",
				phoneNumber: "+7 (123)",
			},
			expectedError:  errBranch,
			expectedResult: pointer.ToInt64(0),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				branchRepository: mocks.NewMockBranchRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewBranchService(
				f.branchRepository,
			)
			id, err := s.AddBranch(tt.args.name, tt.args.address, tt.args.phoneNumber)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, *tt.expectedResult, id)
			}
		})
	}
}

func TestGetAllBranches(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		branchRepository *mocks.MockBranchRepository
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult []*models.Branch
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().GetAllBranches().Return([]*models.Branch{
					{
						ID:          int64(1),
						Name:        "test name",
						Address:     "test addr",
						PhoneNumber: "test phone",
					},
					{
						ID:          int64(2),
						Name:        "test name2",
						Address:     "test addr2",
						PhoneNumber: "test phone2",
					},
				}, nil).Times(1)
			},
			expectedError: nil,
			expectedResult: []*models.Branch{
				{
					ID:          int64(1),
					Name:        "test name",
					Address:     "test addr",
					PhoneNumber: "test phone",
				},
				{
					ID:          int64(2),
					Name:        "test name2",
					Address:     "test addr2",
					PhoneNumber: "test phone2",
				},
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().GetAllBranches().Return(nil, errBranch).Times(1)
			},
			expectedError:  errBranch,
			expectedResult: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				branchRepository: mocks.NewMockBranchRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewBranchService(
				f.branchRepository,
			)
			branches, err := s.GetAllBranches()

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, branches)
			}
		})
	}
}

func TestGetBranchByID(t *testing.T) { //nolint:funlen // это тесты
	type fields struct {
		branchRepository *mocks.MockBranchRepository
	}
	type args struct {
		ID int64
	}

	tests := []struct {
		expectedError  error
		prepare        func(f *fields)
		name           string
		expectedResult models.Branch
		args           args
		wantErr        bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().GetBranchByID(int64(1)).Return(models.Branch{
					ID:          int64(1),
					Name:        "test name",
					Address:     "test addr",
					PhoneNumber: "test phone",
				}, nil).Times(1)
			},
			args: args{
				ID: 1,
			},
			expectedError: nil,
			expectedResult: models.Branch{
				ID:          int64(1),
				Name:        "test name",
				Address:     "test addr",
				PhoneNumber: "test phone",
			},
			wantErr: false,
		},
		{
			name: "err test",
			prepare: func(f *fields) {
				f.branchRepository.EXPECT().GetBranchByID(int64(1)).Return(models.Branch{}, errBranch).Times(1)
			},
			args: args{
				ID: int64(1),
			},
			expectedError:  errBranch,
			expectedResult: models.Branch{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				branchRepository: mocks.NewMockBranchRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := services.NewBranchService(
				f.branchRepository,
			)
			branches, err := s.GetBranchByID(tt.args.ID)

			if tt.wantErr != false {
				assert.Equal(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, branches)
			}
		})
	}
}
