package menu

import "fmt"

//nolint:forbidigo,nolintlint // just why?
func PrintMenuAdmin() {
	fmt.Printf(
		"1) Добавить врача\n" +
			"2) Создать запись\n" +
			"3) Отменить запись\n" +
			"4) Изменить запись\n" +
			"5) Получить запись\n" +
			"6) Получить всех пациентов\n" +
			"7) Получить всех врачей\n")
}

//nolint:forbidigo,nolintlint // just why?
func PrintMenuUser() {
	fmt.Printf(
		"1) Получить мои записи\n" +
			"2) Отменить запись\n" +
			"3) Изменить запись\n" +
			"4) Создать запись\n")
}

//nolint:forbidigo,nolintlint // just why?
func PrintMenuDoctor() {
	fmt.Printf("1) Получить мои записи\n")
}

//nolint:forbidigo,nolintlint // just why?
func PrintMenuAuth() {
	fmt.Printf(
		"1) Войти\n" +
			"2) Зарегистрироваться\n")
}
