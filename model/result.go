package model

import (
	"net/http"
	"strconv"
)

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func (r *Result) Succ(d interface{}) *Result {
	r.Code = strconv.Itoa(http.StatusOK)
	r.Message = ""
	r.Data = d
	r.Success = true
	return r
}

func (r *Result) Fail(code int, message string, d interface{}) *Result {
	r.Code = strconv.Itoa(code)
	r.Message = message
	r.Data = nil
	r.Success = false
	return r
}
