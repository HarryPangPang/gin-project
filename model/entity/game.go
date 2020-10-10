package entity

import "time"

const localDateTimeFormat string = "2006-01-02 15:04:05"

// Game 游戏表
type Game struct {
	ID            uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	GameName      string `gorm:"not null;type:varchar(64)"`
	Alias         string `gorm:"not null;type:varchar(32)"`
	GmID          int64  `gorm:"type:type:int(10)"`
	BbxRegion     string `gorm:"type:varchar(128)"`
	URL           string `gorm:"type:varchar(128)"`
	ImURL         string `gorm:"type:varchar(128)"`
	AccessType    int64  `gorm:"type:int(11)"`
	APISecret     string `gorm:"not null;type:varchar(32)"`
	IsAutoGetData bool
	Sdk           string    `gorm:"type:varchar(128)"`
	Unisdk        string    `gorm:"type:varchar(128)"`
	LastModified  time.Time `gorm:"autoUpdateTime:nano"`
}
