package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yogs696/skilltest/internal/entity"
	"gorm.io/gorm"
)

type (
	ScheduleRepoDB struct {
		db  *gorm.DB
		sql *sql.DB
	}
)

// NewScheduleRepoDB create new DB repo for Schedule entity.
func NewScheduleRepoDB(db *gorm.DB, sql *sql.DB) *ScheduleRepoDB {
	if db != nil {
		return &ScheduleRepoDB{
			db:  db,
			sql: sql,
		}
	}

	return nil
}
func (prd *ScheduleRepoDB) FindByScheduleID(conds map[string]interface{}) (res entity.Schedule, rows int, err error) {
	// Find Schedule records by given conditions and return Schedule.
	tx := prd.db.Find(&res, conds)
	rows = int(tx.RowsAffected)
	err = tx.Error

	return
}

// GetListSchedulePagination select Schedule datas with paginate.
func (prd *ScheduleRepoDB) GetListSchedulePagination(pa *PaginationArgs) (
	countTotal, countFiltered int64,
	res []entity.Schedule,
	err error,
) {

	// [1] Query of count total
	err = prd.db.Model(&entity.Schedule{}).
		Select("COUNT(1) as count").
		Count(&countTotal).
		Error
	if err != nil {
		return
	}

	countFiltered = 0
	if pa.AdditionalFilters != nil {
		// [2] Query of count filtered
		err = prd.db.Model(&entity.Schedule{}).
			Select("COUNT(1) as count").
			Scopes(
				GormDBScope(pa.AdditionalFilters),
			).
			Count(&countFiltered).
			Error
		if err != nil {
			return
		}
	}

	// [3] Query of Schedule datas
	tx := prd.db.Model(&entity.Schedule{}).
		Scopes(
			GormDBScope(pa.AdditionalFilters),
		)

	// [3.3] Set-up query limit
	tx = tx.Limit(int(pa.Limit)).Offset(int(pa.Offset))

	// [4.] Set-up query order
	err = tx.Find(&res).Error
	if err != nil {
		return
	}
	return
}

// insert Schedule
func (prd *ScheduleRepoDB) Insert(a *entity.Schedule) error {
	if a == nil {
		return errors.New("pointer argument cannot be nil")
	}

	// Insert
	if err := prd.db.Omit("updated_at").Create(a).Error; err != nil {
		return err
	}

	return nil
}

// UpdateSchedule by given update statements and conditions.
func (prd *ScheduleRepoDB) UpdateSchedule(
	stats map[string]interface{},
	conds map[string]interface{},
) (int64, error) {
	// Update
	tx := prd.db.Model(&entity.Schedule{}).Where(conds).Updates(stats)

	// Check error
	if tx.Error != nil {
		return 0, tx.Error
	}

	return tx.RowsAffected, nil
}

// delete Schedule
func (prd *ScheduleRepoDB) Delete(scheduleID uint) (row int64, err error) {

	tx := prd.db.Delete(&entity.Schedule{}, scheduleID)
	row = tx.RowsAffected
	err = tx.Error
	fmt.Println("err delete schedule repo:", err)

	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23503") {
			err = fmt.Errorf("schedule id [%v] cannot be deleted", scheduleID)
		}
	}
	return
}
