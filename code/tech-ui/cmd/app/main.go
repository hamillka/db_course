package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"tech-ui/internal/handlers"
	"tech-ui/internal/menu"
)

const (
	ADMIN = iota
	USER
	DOCTOR
)

const (
	LOGIN    = 1
	REGISTER = 2
)

//nolint:forbidigo,nolintlint,funlen,gocognit,cyclop // just why?
func main() {
	menu.PrintMenuAuth()
	var action int64
	_, err := fmt.Scan(&action)
	if err != nil {
		return
	}
	var tokenStr string
	switch action {
	case LOGIN:
		tokenStr = handlers.Login()
	case REGISTER:
		handlers.Register()
	}
	for {
		if action == REGISTER {
			fmt.Printf("\nВХОД В СИСТЕМУ\n")
			tokenStr = handlers.Login()
			action = LOGIN
		}
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("someSecretKey"), nil
		})
		if err != nil {
			return
		}
		newTokenStr := "Bearer " + tokenStr

		switch int64(claims["role"].(float64)) {
		case ADMIN:
			menu.PrintMenuAdmin()
			var adminAction int64
			_, err := fmt.Scan(&adminAction)
			if err != nil {
				fmt.Println("Ошибка ввода, попробуйте снова")
				continue
			}
			switch adminAction {
			case 1:
				handlers.AddDoctor(newTokenStr)
			case 2:
				handlers.AddAppointment(newTokenStr)
			case 3:
				handlers.CancelAppointment(newTokenStr)
			case 4:
				handlers.EditAppointment(newTokenStr)
			case 5:
				handlers.GetAppointmentByID(newTokenStr)
			case 6:
				handlers.GetAllPatients(newTokenStr)
			case 7:
				handlers.GetAllDoctors(newTokenStr)
			}
		case USER:
			menu.PrintMenuUser()
			var userAction int64
			_, err := fmt.Scan(&userAction)
			if err != nil {
				fmt.Println("Ошибка ввода, попробуйте снова")
				continue
			}
			switch userAction {
			case 1:
				handlers.GetAppointmentsByPatient(newTokenStr)
			case 2:
				handlers.CancelAppointmentPatient(newTokenStr)
			case 3:
				handlers.EditAppointmentPatient(newTokenStr)
			case 4:
				handlers.AddAppointmentPatient(newTokenStr)
			}
		case DOCTOR:
			menu.PrintMenuDoctor()
			var doctorAction int64
			_, err := fmt.Scan(&doctorAction)
			if err != nil {
				return
			}
			if doctorAction == 1 {
				handlers.GetAppointmentByDoctor(newTokenStr)
			}
		default:
			break
		}
	}
}
