package models

type Request struct {
	Method    string
	Route     string
	Body      string
	RespBody  interface{}
	Headers   [][2]string
	ParseResp bool
}
