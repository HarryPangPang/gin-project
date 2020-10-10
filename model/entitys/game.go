package entitys

import "time"

const localDateTimeFormat string = "2006-01-02 15:04:05"

// Game 游戏表
type Game struct {
	ID            int64     `xorm:"not null INT(10) pk autoincr"`
	GameName      string    `xorm:"not null VARCHAR(64)"`
	Alias         string    `xorm:"not null VARCHAR(32)"`
	GmID          int64     `xorm:"INT(10)"`
	BbxRegion     string    `xorm:"VARCHAR(128)"`
	URL           string    `xorm:"VARCHAR(128)"`
	ImURL         string    `xorm:"VARCHAR(128)"`
	AccessType    int       `xorm:"INT(11)"`
	APISecret     string    `xorm:"not null VARCHAR(32)"`
	IsAutoGetData bool      `xorm:"BOOL"`
	Sdk           string    `xorm:"VARCHAR(128)"`
	Unisdk        string    `xorm:"VARCHAR(128)"`
	LastModified  time.Time `xorm:"updated not null"`
}
