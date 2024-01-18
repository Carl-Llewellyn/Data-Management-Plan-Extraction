package jsonmodels

import "github.com/google/uuid"

type SessionRecords struct {
	Data []SessionRecord `json:"data"`
}

type SessionRecord struct {
	Uuid     uuid.UUID `gorm:"column:uuid" json:"uuid"`
	Username string    `gorm:"column:username" json:"username"`
}

func (SessionRecord) TableName() string {
	return "ftf.session_records"
}
