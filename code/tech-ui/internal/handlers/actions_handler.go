package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tech-ui/internal/dto"
	"tech-ui/internal/models"
)

func CreateRequest(method, route, body string, headers [][2]string) (*http.Request, error) {
	//nolint: noctx //ctx is for what?
	req, err := http.NewRequest(method, "http://localhost:8080"+route, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for idx := range headers {
		req.Header.Add(headers[idx][0], headers[idx][1])
	}

	return req, nil
}

func SendRequest(
	client *http.Client,
	req *http.Request,
	parseResp bool,
	respBody interface{},
) {
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if parseResp {
		err := json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return
		}
	}

	resp.Body.Close()
}

//nolint:forbidigo,nolintlint // just why?
func AddDoctor(token string) {
	var fio, email, phoneNumber string
	fmt.Printf("Введите ФИО: ")
	_, err := fmt.Scan(&fio)
	if err != nil {
		return
	}
	fmt.Printf("Введите email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		return
	}
	fmt.Printf("Введите номер телефона: ")
	_, err = fmt.Scan(&phoneNumber)
	if err != nil {
		return
	}

	resp := dto.CreateOrEditDoctorResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method: "POST",
		Route:  "/api/v1/doctor",
		Body: fmt.Sprintf("{\"fio\": \"%v\", \"email\": \"%v\", \"phoneNumber\": \"%v\"}",
			fio, email, phoneNumber),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("%v", resp.ID)
}

//nolint:forbidigo,nolintlint // just why?
func AddAppointment(token string) {
	var (
		patientID, doctorID int64
		dateTime            string
	)
	fmt.Printf("Введите id пациента: ")
	_, err := fmt.Scan(&patientID)
	if err != nil {
		return
	}
	fmt.Printf("Введите id доктора: ")
	_, err = fmt.Scan(&doctorID)
	if err != nil {
		return
	}
	fmt.Printf("Введите дату: ")
	_, err = fmt.Scan(&dateTime)
	if err != nil {
		return
	}

	resp := dto.CreateOrEditAppointmentResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method: "POST",
		Route:  "/api/v1/appointment",
		Body: fmt.Sprintf("{\"patientId\": %v, \"doctorId\": %v, \"dateTime\": \"%v\"}",
			patientID, doctorID, dateTime),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("%v\n", resp)
}

//nolint:forbidigo,nolintlint // just why?
func CancelAppointment(token string) {
	var id int64
	fmt.Printf("Введите id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "DELETE",
		Route:     fmt.Sprintf("/api/v1/appointment/%v", id),
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: false,
		RespBody:  nil,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)
}

//nolint:forbidigo,nolintlint // just why?
func EditAppointment(token string) {
	var (
		id, patientID, doctorID int64
		dateTime                string
	)
	fmt.Printf("Введите id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}
	fmt.Printf("Введите id пациента: ")
	_, err = fmt.Scan(&patientID)
	if err != nil {
		return
	}
	fmt.Printf("Введите id доктора: ")
	_, err = fmt.Scan(&doctorID)
	if err != nil {
		return
	}
	fmt.Printf("Введите дату: ")
	_, err = fmt.Scan(&dateTime)
	if err != nil {
		return
	}

	resp := dto.CreateOrEditAppointmentResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method: "PATCH",
		Route: fmt.Sprintf("/api/v1/appointment/%v",
			id),
		Body: fmt.Sprintf("{\"patientId\": %v, \"doctorId\": %v, \"dateTime\": \"%v\"}",
			patientID, doctorID, dateTime),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("%v\n", resp.ID)
}

//nolint:forbidigo,nolintlint // just why?
func GetAppointmentByID(token string) {
	var id int64
	fmt.Printf("Введите id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	resp := dto.GetAppointmentResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "GET",
		Route:     fmt.Sprintf("/api/v1/appointment/%v", id),
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("id: %v doctor: %v patient: %v date: %v\n",
		resp.ID, resp.DoctorID, resp.PatientID, resp.DateTime)
}

//nolint:forbidigo,nolintlint // just why?
func GetAllPatients(token string) {
	var resp []dto.GetPatientResponseDto
	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "GET",
		Route:     "/api/v1/patient",
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	for idx := range resp {
		fmt.Printf("%v %v %v %v %v\n", resp[idx].ID, resp[idx].Fio,
			resp[idx].Email, resp[idx].PhoneNumber, resp[idx].Insurance)
	}
}

//nolint:forbidigo,nolintlint // just why?
func GetAllDoctors(token string) {
	var resp []dto.GetDoctorResponseDto
	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "GET",
		Route:     "/api/v1/doctor",
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	for idx := range resp {
		fmt.Printf("%v %v %v %v\n", resp[idx].ID, resp[idx].Fio, resp[idx].Email, resp[idx].PhoneNumber)
	}
}

// GetAppointmentByDoctor for doctors
//
//nolint:forbidigo,nolintlint // just why?
func GetAppointmentByDoctor(token string) {
	fmt.Printf("Введите id: ")
	var id int64
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	var resp []dto.GetAppointmentResponseDto

	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "GET",
		Route:     fmt.Sprintf("/api/v1/appointment?doctor_id=%v", id),
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	for idx := range resp {
		fmt.Printf("%v %v %v %v\n",
			resp[idx].ID, resp[idx].PatientID, resp[idx].DoctorID, resp[idx].DateTime)
	}
}

// for patients
//
//nolint:forbidigo,nolintlint // just why?
func GetAppointmentsByPatient(token string) {
	fmt.Printf("Введите id: ")
	var id int64
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	var resp []dto.GetAppointmentResponseDto

	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "GET",
		Route:     fmt.Sprintf("/api/v1/appointment?patient_id=%v", id),
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	for idx := range resp {
		fmt.Printf("%v %v %v %v\n",
			resp[idx].ID, resp[idx].PatientID, resp[idx].DoctorID, resp[idx].DateTime)
	}
}

//nolint:forbidigo,nolintlint // just why?
func CancelAppointmentPatient(token string) {
	fmt.Printf("Введите id: ")
	var id int64
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	client := http.Client{}
	reqInfo := &models.Request{
		Method:    "DELETE",
		Route:     fmt.Sprintf("/api/v1/appointment/%v", id),
		Body:      "",
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: false,
		RespBody:  nil,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)
}

//nolint:forbidigo,nolintlint // just why?
func EditAppointmentPatient(token string) {
	var (
		id, patientID, doctorID int64
		dateTime                string
	)
	fmt.Printf("Введите id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}
	fmt.Printf("Введите id пациента: ")
	_, err = fmt.Scan(&patientID)
	if err != nil {
		return
	}
	fmt.Printf("Введите id доктора: ")
	_, err = fmt.Scan(&doctorID)
	if err != nil {
		return
	}
	fmt.Printf("Введите дату: ")
	_, err = fmt.Scan(&dateTime)
	if err != nil {
		return
	}

	resp := dto.CreateOrEditAppointmentResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method: "PATCH",
		Route: fmt.Sprintf("/api/v1/appointment/%v",
			id),
		Body: fmt.Sprintf("{\"patientId\": %v, \"doctorId\": %v, \"dateTime\": \"%v\"}",
			patientID, doctorID, dateTime),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("%v\n", resp.ID)
}

//nolint:forbidigo,nolintlint // just why?
func AddAppointmentPatient(token string) {
	var (
		patientID, doctorID int64
		dateTime            string
	)
	fmt.Printf("Введите id пациента: ")
	_, err := fmt.Scan(&patientID)
	if err != nil {
		return
	}
	fmt.Printf("Введите id доктора: ")
	_, err = fmt.Scan(&doctorID)
	if err != nil {
		return
	}
	fmt.Printf("Введите дату: ")
	_, err = fmt.Scan(&dateTime)
	if err != nil {
		return
	}

	resp := dto.CreateOrEditAppointmentResponseDto{}
	client := http.Client{}
	reqInfo := &models.Request{
		Method: "POST",
		Route:  "/api/v1/appointment",
		Body: fmt.Sprintf("{\"patientId\": %v, \"doctorId\": %v, \"dateTime\": \"%v\"}",
			patientID, doctorID, dateTime),
		Headers:   [][2]string{{"Content-Type", "application/json"}},
		ParseResp: true,
		RespBody:  &resp,
	}
	req, err := CreateRequest(
		reqInfo.Method,
		reqInfo.Route,
		reqInfo.Body,
		append(reqInfo.Headers, [2]string{"auth-x", token}),
	)
	if err != nil {
		return
	}
	SendRequest(&client, req, reqInfo.ParseResp, reqInfo.RespBody)

	fmt.Printf("%v\n", resp)
}
