package model

import (
	"net/http"
	"strconv"
	)

type ResponseModel struct {
	Code string		    `json:"code"`
	Message string		`json:"message"`
	Data interface{}    `json:"data"`
	Success bool        `json:"success"`
}

func (r *ResponseModel) Succ(d interface{}) *ResponseModel{
	r.Code = strconv.Itoa(http.StatusOK)
	r.Message = ""
	r.Data = d
	r.Success = true
	return r
}

