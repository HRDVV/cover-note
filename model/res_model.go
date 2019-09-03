package model

import (
	"net/http"
	"strconv"
)

type ResponseModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func (r *ResponseModel) Succ(d interface{}) *ResponseModel {
	r.Code = strconv.Itoa(http.StatusOK)
	r.Message = ""
	r.Data = d
	r.Success = true
	return r
}

func (r *ResponseModel) Fail(code string, message string, d interface{}) *ResponseModel {
	r.Code = code
	r.Message = message
	r.Data = nil
	r.Success = false
	return r
}
