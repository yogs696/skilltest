package schedule

import (
	"errors"
	"fmt"
	"time"

	"github.com/yogs696/skilltest/internal/entity"
	"github.com/yogs696/skilltest/internal/repo"
	"github.com/yogs696/skilltest/pkg/kemu"
)

// Service represent Schedule services interface
type (
	Service struct {
		repo Repository
		kemu *kemu.Mutex
	}

	PSchedule repo.PaginationArgs
)

// NewService creates new Schedule services
func NewService(
	kemu *kemu.Mutex,
	r Repository,
	callback ...func(s string),
) *Service {
	if len(callback) > 0 {
		callback[0]("Registering Schedule List Domain Entity...")
	}

	svc := &Service{
		repo: r,
		kemu: kemu,
	}

	return svc
}

// CreateSchedule create new schedule
func (s *Service) CreateSchedule(
	cinemaId, movieId uint,
	showDate time.Time,
	startTime, endTime string,
) error {

	a := &entity.Schedule{
		CinemaID:  cinemaId,
		MovieID:   movieId,
		ShowDate:  showDate,
		StartTime: startTime,
		EndTime:   endTime,
	}

	defer func() {
		a = nil
	}()

	if err := s.repo.Insert(a); err != nil {
		return err
	}

	return nil
}

// Createschedule create new schedule
func (s *Service) UpdateSchedule(
	scheduleId uint,
	cinemaId, movieId uint,
	showDate time.Time,
	startTime, endTime string,
) (int64, error) {
	stmt := map[string]interface{}{
		"cinema_id":  cinemaId,
		"movie_id":   movieId,
		"show_date":  showDate,
		"start_time": startTime,
		"end_time":   endTime,
	}

	cond := map[string]interface{}{"id": scheduleId}

	return s.repo.UpdateSchedule(stmt, cond)
}

// GetScheduleListPagination Schedule's GetListPagination wrapper.
func (s *Service) GetScheduleListPagination(pp *PSchedule) (
	countTotal, countFiltered int64,
	res []entity.Schedule,
	err error,
) {
	if pp == nil {
		err = errors.New("arguments cannot be nil")
		return
	}

	total, filtered, x, err := s.repo.GetListSchedulePagination((*repo.PaginationArgs)(pp))
	if err != nil {
		return
	}

	countTotal = total
	countFiltered = filtered
	res = x
	return
}

// FindByScheduleId
func (s *Service) FindByScheduleId(scheduleId interface{}) (*entity.Schedule, int, error) {
	// Prepare query conditions
	conds := make(map[string]interface{})
	switch {
	case scheduleId != nil:
		conds["id"] = scheduleId

	default:
		return nil, 0, fmt.Errorf("argument scheduleId cannot be empty -> scheduleId: %d", scheduleId)
	}

	schedule, rows, err := s.repo.FindByScheduleID(conds)

	switch {
	// Always check error appear first
	case err != nil:
		return nil, rows, err

	case rows == 0:
		return nil, rows, nil

	default:
		return &schedule, rows, nil
	}
}

// delete Schedule by id
func (s *Service) DeleteByScheduleId(scheduleId uint) (row int64, err error) {

	return s.repo.Delete(scheduleId)
}
