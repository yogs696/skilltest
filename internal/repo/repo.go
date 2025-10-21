package repo

import (
	"github.com/yogs696/skilltest/internal/entity/std"
	"gorm.io/gorm"
)

type (
	AdditionalConditions struct {
		Qry interface{}
		Arg []interface{}

		GroupingOr []AdditionalConditions
	}

	// CustomUpdateStatements hold data structtur for updating record using custom statement.
	CustomUpdateStatements struct {
		UsingExpression bool
		Column          string
		Expr            string
		Statement       []interface{}
	}

	// PaginationArgs hold type of data pagination arguments.
	PaginationArgs struct {
		AdditionalFilters map[interface{}][]interface{}
		Joins             map[interface{}][]interface{}
		Limit             int64
		Offset            int64
	}

	// RepoDatatableResponse defines datatable response
	RepoDatatableResponse std.ResponseDatatable
)

// GormDBScope make scoped chained DB query and return it as a gorm DB pointer.
func GormDBScope(s map[interface{}][]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for q, a := range s {
			if len(a) > 0 {
				db = db.Where(q, a...)
			} else {
				db = db.Where(q)
			}
		}

		return db
	}
}
