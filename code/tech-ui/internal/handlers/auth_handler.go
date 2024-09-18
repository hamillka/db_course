package handlers

import (
	"fmt"
	"net/http"
	"tech-ui/internal/dto"
	"tech-ui/internal/models"
)

//nolint:forbidigo,nolintlint // just why?
func Login() string {
	var login, password string
	fmt.Printf("Введите логин: ")
	_, err := fmt.Scan(&login)
	if err != nil {
		return ""
	}
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return ""
	}

	resp := dto.UserLoginResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "POST",
		Route:     "/auth/login",
		Body:      fmt.Sprintf("{\"login\": \"%v\", \"password\": \"%v\"}", login, password),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(reqInfo.Method, reqInfo.Route, reqInfo.Body, reqInfo.Headers)
	if err != nil {
		return ""
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	return resp.JWTToken
}

//nolint:forbidigo,nolintlint // just why?
func Register() {
	var (
		fio, phoneNumber, email, login, password, insurance string
		role                                                int64
	)
	fmt.Printf("Введите ФИО: ")
	_, err := fmt.Scan(&fio)
	if err != nil {
		return
	}
	fmt.Printf("Введите номер телефона: ")
	_, err = fmt.Scan(&phoneNumber)
	if err != nil {
		return
	}
	fmt.Printf("Введите email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		return
	}
	fmt.Printf("Введите логин: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return
	}
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return
	}
	fmt.Printf("Введите роль: ")
	_, err = fmt.Scan(&role)
	if err != nil {
		return
	}
	if role == dto.USER {
		fmt.Printf("Введите номер страховки: ")
		_, err = fmt.Scan(&insurance)
		if err != nil {
			return
		}
	}

	client := http.Client{}
	reqInfo := &models.Request{
		Method: "POST",
		Route:  "/auth/register",
		Body: fmt.Sprintf("{\"fio\": \"%v\", \"phoneNumber\": \"%v\", "+
			"\"email\": \"%v\", \"insurance\": \"%v\", \"role\": %v, "+
			"\"login\": \"%v\", \"password\": \"%v\"}",
			fio, phoneNumber, email, insurance, role, login, password),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: false,
		RespBody:  nil,
	}
	req, err := CreateRequest(reqInfo.Method, reqInfo.Route, reqInfo.Body, reqInfo.Headers)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)
}
