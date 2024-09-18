package dto

type CreateOrEditPatientRequestDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Insurance   string `json:"insurance"`
	ID          int64  `json:"id"`
}

type CreateOrEditPatientResponseDto struct {
	ID int64 `json:"id"`
}

type GetPatientResponseDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Insurance   string `json:"insurance"`
	ID          int64  `json:"id"`
}
