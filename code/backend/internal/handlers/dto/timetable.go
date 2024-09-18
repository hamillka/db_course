package dto

import "github.com/hamillka/ppo/backend/internal/models"

type GetTimetableResponseDto struct {
	WorkDay  int64 `json:"workDay"`
	ID       int64 `json:"id"`
	DoctorID int64 `json:"doctorId"`
	OfficeID int64 `json:"officeId"`
}

func ConvertToTimetableDto(tt *models.Timetable) *GetTimetableResponseDto {
	return &GetTimetableResponseDto{
		WorkDay:  tt.WorkDay,
		ID:       tt.ID,
		DoctorID: tt.DoctorID,
		OfficeID: tt.OfficeID,
	}
}
