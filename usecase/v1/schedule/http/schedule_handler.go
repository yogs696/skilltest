package uv1schedulehttp

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/internal/entity/std"
	"github.com/yogs696/skilltest/internal/helper"
	"github.com/yogs696/skilltest/internal/repo"
	"github.com/yogs696/skilltest/usecase/v1/schedule"
)

func (ds *domainService) Testing(c echo.Context) error {
	var apiResp *std.APIResponse

	data := helper.CtxValue(c)
	apiResp = std.APIResponseSuccess(data)
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// create schedule
func (ds *domainService) List(c echo.Context) error {
	p := new(list__Reqp)

	var (
		apiResp           *std.APIResponse
		AdditionalFilters map[interface{}][]interface{}
	)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Prepare response
	resp := &repo.RepoDatatableResponse{Draw: p.Draw + 1}
	defer func() {
		resp = nil
	}()

	if len(p.Search) > 0 {
		sID64, err := strconv.ParseUint(p.Search, 10, 64) // base 10, bit size 64
		if err != nil {
			log.Fatalf("Error parsing string to uint16: %v", err)
			return c.JSON(int(std.APIStatusBadRequest), err)
		}

		// Since ParseUint returns a uint, we need to explicitly cast it to uint
		sIDUint := uint(sID64)

		AdditionalFilters = map[interface{}][]interface{}{
			"cinema_id =? OR movie_id =?": {sIDUint, sIDUint},
		}
	}

	pp := &schedule.PSchedule{
		AdditionalFilters: AdditionalFilters,
		Limit:             int64(p.Length),
		Offset:            int64(p.Offset),
	}

	countTotoal, countfiltered, list, err := ds.s.GetScheduleListPagination(pp)
	if err != nil {
		log.Printf("[ListSchedule] - Error: %s", err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	resp.RecordsTotal = countTotoal
	resp.FilteredTotal = countfiltered
	resp.Data = list

	// Return success
	return c.JSON(int(std.StatusOK), resp)
}

// create schedule
func (ds *domainService) Create(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(schedule__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Log payload
	log.Printf("[CreateSchedule][Payload] - %v", *p)

	showDate, err := time.Parse("2006-01-02 15:04:05", p.ShowDate)
	if err != nil {
		return c.JSON(int(std.StatusBadRequest), err)
	}

	if err := ds.s.CreateSchedule(p.CinemaId, p.MovieId, showDate, p.StartTime, p.EndTime); err != nil {
		log.Printf("[CreateSchedule] - Error: %s", err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess("Data Schedule successfully created")
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// Update schedule
func (ds *domainService) Update(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(schedule__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	sID64, err := strconv.ParseUint(c.Param("id"), 10, 64) // base 10, bit size 64
	if err != nil {
		log.Fatalf("Error parsing string to uint16: %v", err)
	}

	// Since ParseUint returns a uint, we need to explicitly cast it to uint
	sIDUint := uint(sID64)

	// Log payload
	log.Printf("[UpdateSchedule-sID[%d]][Payload] - %v", sIDUint, *p)

	// Update record
	showDate, err := time.Parse("2006-01-02 15:04:05", p.ShowDate)
	if err != nil {
		return c.JSON(int(std.StatusBadRequest), err)
	}

	_, err = ds.s.UpdateSchedule(sIDUint, p.CinemaId, p.MovieId, showDate, p.StartTime, p.EndTime)
	if err != nil {
		log.Printf("[UpdateSchedule] - Error: %s", err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess("Data Schedule successfully updated")
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// delete Schedule
func (ds *domainService) Delete(c echo.Context) error {
	var apiResp *std.APIResponse

	sID64, err := strconv.ParseUint(c.Param("id"), 10, 64) // base 10, bit size 64
	if err != nil {
		log.Fatalf("Error parsing string to uint16: %v", err)
	}

	// Since ParseUint returns a uint, we need to explicitly cast it to uint
	sIDUint := uint(sID64)

	rows, err := ds.s.DeleteByScheduleId(sIDUint)

	switch {
	// Always check error appear first
	case err != nil:
		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}

	case rows == 0:
		if apiResp = std.APIResponseError(std.StatusNotFound, errors.New("record not found")); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess("Data Schedule successfully deleted")
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
