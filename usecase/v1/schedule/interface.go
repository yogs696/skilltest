package schedule

import (
	"github.com/yogs696/skilltest/internal/entity"
	"github.com/yogs696/skilltest/internal/repo"
)

type Repository interface {
	GetListSchedulePagination(pa *repo.PaginationArgs) (
		countTotal, countFiltered int64,
		res []entity.Schedule,
		err error,
	)
	FindByScheduleID(conds map[string]interface{}) (res entity.Schedule, row int, err error)
	UpdateSchedule(
		stats map[string]interface{},
		conds map[string]interface{},
	) (int64, error)

	Insert(a *entity.Schedule) error
	Delete(ScheduleID uint) (row int64, err error)
}
