package integration_test

import (
	"net/http"
	"testing"

	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	net "github.com/hamillka/ppo/backend/internal/integration_test"
	"github.com/stretchr/testify/require"
)

var resp dto.UserLoginResponseDto //nolint:gochecknoglobals // для передачи токена между тестами

func TestCreateAppointment(t *testing.T) {
	client := http.Client{}

	tests := []struct {
		respBody       interface{}
		Name           string
		Method         string
		route          string
		body           string
		headers        [][2]string
		expectedStatus int
		parseResp      bool
	}{
		{
			Name:   "user register success",
			Method: http.MethodPost,
			route:  "/auth/register",
			body: "{\"fio\": \"a\",\"phoneNumber\": \"b\", " +
				"\"email\": \"c\",\"insurance\": \"d\", \"login\": \"e\", " +
				"\"password\": \"f\", \"role\": 1}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusOK,
			parseResp:      false,
			respBody:       nil,
		},
		{
			Name:           "user login success",
			Method:         http.MethodPost,
			route:          "/auth/login",
			body:           "{\"login\": \"e\", \"password\": \"f\"}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusOK,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:   "user appointment create success",
			Method: http.MethodPost,
			route:  "/api/v1/appointment",
			body: "{\"dateTime\": \"2024-03-15T10:30:00Z\", " +
				"\"patientId\": 100, " +
				"\"doctorId\": 200}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusCreated,
			parseResp:      false,
			respBody:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var (
				req *http.Request
				err error
			)
			if tt.Name == "user appointment create success" {
				req, err = net.CreateRequest(tt.Method, tt.route, tt.body,
					append(tt.headers, [2]string{"auth-x", "Bearer " + resp.JWTToken}))
			} else {
				req, err = net.CreateRequest(tt.Method, tt.route, tt.body, tt.headers)
			}
			require.NoError(t, err)
			net.SendRequest(t, &client, req, tt.expectedStatus, tt.parseResp, tt.respBody)
		})
	}
}
