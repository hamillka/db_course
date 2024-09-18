package dto

type CreateOrEditDoctorRequestDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	ID          int64  `json:"id"`
}

type CreateOrEditDoctorResponseDto struct {
	ID int64 `json:"id"`
}

type GetDoctorResponseDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	ID          int64  `json:"id"`
}
