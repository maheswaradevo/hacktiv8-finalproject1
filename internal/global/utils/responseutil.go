package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject1/pkg/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/pkg/errors"
)

func NewSuccessResponsWriter(rw http.ResponseWriter, code int, status string, data interface{}) {
	BaseResponseWriter(rw, code, status, nil, data)
}

func NewErrorResponse(rw http.ResponseWriter, err error) {
	errMap := errors.GetErrorResponseMetaData(err)
	BaseResponseWriter(rw, errMap.Code, "", &dto.ErrorData{Code: errMap.Code, Message: errMap.Message}, nil)
}

func BaseResponseWriter(rw http.ResponseWriter, code int, status string, er *dto.ErrorData, data interface{}) {
	res := dto.BaseResponse{
		Status: status,
		Data:   data,
		Error:  er,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Printf("cant marshal the interface")
	}
	rw.Header().Add("Content Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(jsonData)
}
