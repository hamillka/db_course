package dto

const (
	ADMIN = iota
	USER
	DOCTOR
)

type UserLoginRequestDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLoginResponseDto struct {
	JWTToken string `json:"jwtToken"`
}

type UserRegisterRequestDto struct {
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Insurance   string `json:"insurance"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Role        int64  `json:"role"`
}
