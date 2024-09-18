package integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	t *testing.T,
	client *http.Client,
	req *http.Request,
	expectedStatus int,
	parseResp bool,
	respBody interface{},
) {
	t.Helper()

	resp, err := client.Do(req)
	require.NoError(t, err)

	require.Equal(t, expectedStatus, resp.StatusCode)

	if parseResp {
		err := json.NewDecoder(resp.Body).Decode(&respBody)
		require.NoError(t, err)
	}

	resp.Body.Close()
}
