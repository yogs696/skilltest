package helper

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/config"
	"github.com/yogs696/skilltest/internal/entity/std"
)

func InArray(v interface{}, in interface{}) (ok bool) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

// Will return date in date format 'YYYY-mm-DD' if DetTime is valid.
func DateStr() (d string) {
	loc, err := time.LoadLocation(config.Of.App.TimeZone)
	if err != nil {
		fmt.Println("PANIC_LOGIC")
		panic(err)
	}

	d = time.Now().In(loc).Format("2006-01-02")
	return
}

func LoadLocation() time.Time {
	loc, err := time.LoadLocation(config.Of.App.TimeZone)
	if err != nil {
		fmt.Println("PANIC_LOGIC")
		panic(err)
	}

	return time.Now().In(loc)
}

func StructValidator(c echo.Context, p interface{}) *std.APIResponse {
	// Parsing params/payload to struct
	if err := c.Bind(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, errors.New("failed to parsing request params/payloads"))
	}

	// Validate the params/payloads
	if err := c.Validate(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, err)
	}

	return nil
}
