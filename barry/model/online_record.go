package model

import "time"

type OnlineRecord struct {
	ID uint `gorm:"primaryKey"`
	Project string `gorm:"column:project"`
	Date time.Time `gorm:"column: date"`
	OnlineContent string `gorm:"column:online_content"`
	TapdLink string `gorm:"column:tapd_link"`
	FileType string `gorm:"column:file_type"`
	Initiator string `gorm:"column:init_type"`
	OnlineType string `gorm:"column:online_type"`
	Reason string `gorm:"reason"`
	Detail string `gorm:"detail"`
	Way string `gorm:"way"`
	Mint uint `gorm:"mint"`
}

func (OnlineRecord) TableName() string {
	return "online_record"
}