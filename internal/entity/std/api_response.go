// Package std is represent standard from entoty domain.
// Can be following standardization of the user
package std

import (
	"fmt"

	"github.com/yogs696/skilltest/usecase"
)

type (
	// APIResponse defines default RESTful API response standard
	APIResponse struct {
		StatusCode HTTPStatusCode
		Body       Response
	}

	// Response defines default RESTful API body
	Response struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	}

	HTTPStatusCode int
	APIStatusCode  int

	ResponseDatatable struct {
		Draw          int         `json:"draw"`
		RecordsTotal  int64       `json:"recordsTotal"`
		FilteredTotal int64       `json:"filteredTotal"`
		Data          interface{} `json:"data"`
	}
)

// Standard HTTP status
const (
	// HTTP/1.1 2xx
	StatusOK HTTPStatusCode = 200 // RFC 7231, 6.3.1

	// HTTP/1.1 4xx
	StatusBadRequest          HTTPStatusCode = 400 // RFC 7231, 6.5.1
	StatusForbidden           HTTPStatusCode = 403 // RFC 7231, 6.5.3
	StatusNotFound            HTTPStatusCode = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed    HTTPStatusCode = 405 // RFC 7231, 6.5.5
	StatusUnprocessableEntity HTTPStatusCode = 422 // RFC 4918, 11.2

	// HTTP/1.1 5xx
	StatusServerError        HTTPStatusCode = 500 // RFC 7231, 6.6.1
	StatusServiceUnavailable HTTPStatusCode = 503 // RFC 7231, 6.6.4
)

// Standard API status
const (
	APIStatusOK APIStatusCode = 2400

	APIStatusBadRequest          APIStatusCode = 2401
	APIStatusServerError         APIStatusCode = 2402
	APIStatusUnprocessableEntity APIStatusCode = 2403
	APIStatusServiceUnavailable  APIStatusCode = 2404
	APIStatusNotFound            APIStatusCode = 2405
	APIStatusForbidden           APIStatusCode = 2406

	// Warning code
	TOOMANYREQUESTS APIStatusCode = 2410
)

// Variable mapping from http status code to API status code
var httpStatus2APIStatus = map[HTTPStatusCode]APIStatusCode{
	StatusOK: APIStatusOK,

	StatusBadRequest:          APIStatusBadRequest,
	StatusForbidden:           APIStatusForbidden,
	StatusNotFound:            APIStatusNotFound,
	StatusMethodNotAllowed:    APIStatusServiceUnavailable,
	StatusUnprocessableEntity: APIStatusUnprocessableEntity,

	StatusServerError:        APIStatusServerError,
	StatusServiceUnavailable: APIStatusServiceUnavailable,
}

// APIErrorResponse return standard API error response
func APIResponseError(sc HTTPStatusCode, err error, customResCode ...APIStatusCode) *APIResponse {
	var data interface{} = nil
	resCode := int(httpStatus2APIStatus[sc])
	if len(customResCode) > 0 {
		resCode = int(customResCode[0])
	}

	if err != nil {
		fmt.Printf("APIResponseError - Error %s \n", err.Error())
		return &APIResponse{
			StatusCode: sc,
			Body: Response{
				Success: false,
				Code:    resCode,
				Data:    data,
				Error:   usecase.ParseUnwantedError(err),
			},
		}
	}

	return nil
}

// APIResponseSuccess return standard API success response
func APIResponseSuccess(d interface{}) *APIResponse {
	sc := 200

	return &APIResponse{
		StatusCode: HTTPStatusCode(sc),
		Body: Response{
			Success: true,
			Code:    int(httpStatus2APIStatus[HTTPStatusCode(sc)]),
			Data:    d,
			Error:   nil,
		},
	}
}
