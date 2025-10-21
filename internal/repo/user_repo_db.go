package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/yogs696/skilltest/internal/entity"
	"gorm.io/gorm"
)

type UserRepoDB struct {
	db  *gorm.DB
	sql *sql.DB
}

// NewUserRepoDB create new DB repo for User entity.
func NewUserRepoDB(db *gorm.DB, sql *sql.DB) *UserRepoDB {
	if db != nil {
		return &UserRepoDB{
			db:  db,
			sql: sql,
		}
	}

	return nil
}

// Transaction repo method of User that approach DB process with transaction.
func (urd *UserRepoDB) Transaction(txFunc func(interface{}) error) (err error) {
	tx := urd.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if err != nil {
			log.Printf("[DBTxSession] - Rollback with reason: %s", err.Error())
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				log.Printf("[DBTxSession] - Commit error: %s", err.Error())
			}
		}
	}()

	err = txFunc(tx)
	return
}

// Create new User record.
func (urd *UserRepoDB) Create(w *entity.User) (*entity.User, error) {
	// w should != nil
	if w == nil {
		return nil, errors.New("entity user argument pointer cannot be nil")
	}

	// Manually freed memory
	defer func() {
		w = nil
	}()

	// Insert
	if err := urd.db.Create(w).Error; err != nil {
		switch {
		case strings.Contains(err.Error(), "SQLSTATE 23505"):
			err = fmt.Errorf("user with email : %s is already exists", w.Email)
		}

		return nil, err
	}

	return w, nil
}

// Create new User record.
func (urd *UserRepoDB) FindByEamil(conds map[string]interface{}) (res entity.User, rows int, err error) {
	// Query
	tx := urd.db.Find(&res, conds)
	rows = int(tx.RowsAffected)
	err = tx.Error
	return
}
