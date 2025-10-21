package uv1authroute

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/internal/entity/std"
	"github.com/yogs696/skilltest/internal/helper"
	"golang.org/x/crypto/bcrypt"
)

// create register
func (ds *domainService) Register(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(register__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Log payload
	log.Printf("[RegisterUser-%s][Payload] - %v", p.Username, *p)

	usr, err := ds.s.CreateUser(p.Username, p.Email, p.Password)
	if err != nil {
		log.Printf("[RegisterUser] - (%s) Error: %s", p.Username, err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}
	// generate token
	token, err := helper.GenerateToken(usr.ID, p.Username, p.Email)
	if err != nil {
		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	formatter := helper.FormatUser(usr, token)
	apiResp = std.APIResponseSuccess(formatter)
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

func (ds *domainService) Login(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(Login__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Log payload
	log.Printf("[LoginUser-%s][Payload] - %v", p.Email, *p)

	// compare passowrd

	usr, rows, err := ds.s.LoginUser(p.Email, p.Password)

	switch {
	case err != nil || rows == 0:
		log.Printf("[LoginUser] - (%s) Error: %s", p.Email, err.Error())

		statusErr := std.StatusServerError
		if rows == 0 {
			statusErr = std.StatusNotFound
		}

		if apiResp = std.APIResponseError(statusErr, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}

	default:
		// check password
		err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(p.Password))
		if err != nil {
			if apiResp = std.APIResponseError(std.StatusBadRequest, errors.New("Email or password not valid")); apiResp != nil {
				return c.JSON(int(apiResp.StatusCode), apiResp.Body)
			} else {
				return c.NoContent(500)
			}
		}
	}

	// generate token
	token, err := helper.GenerateToken(usr.ID, p.Email, p.Email)
	if err != nil {
		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	formatter := helper.FormatUser(usr, token)
	apiResp = std.APIResponseSuccess(formatter)
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
