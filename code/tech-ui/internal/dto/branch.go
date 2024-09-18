package dto

type CreateOrEditBranchRequestDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	ID          int64  `json:"id"`
}

type CreateOrEditBranchResponseDto struct {
	ID int64 `json:"id"`
}

type GetBranchResponseDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	ID          int64  `json:"id"`
}
