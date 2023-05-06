package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Pagination struct {
	Current  int   `json:"-"`
	PageSize int   `json:"-"`
	Total    int64 `json:"-"`
}

func (t Pagination) Sql() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if t.Current <= 0 {
			t.Current = 1
		}
		switch {
		case t.PageSize < 6:
			t.PageSize = 5
		case t.PageSize > 5:
			t.PageSize = 10
		}
		return db.Offset((t.Current - 1) * t.PageSize).Limit(t.PageSize)
	}
}
