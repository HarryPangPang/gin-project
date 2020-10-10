package entity

import "time"

// ChannelAppIds 游戏和channelappId对应表
type ChannelAppIds struct {
	ID           uint      `gorm:"primaryKey;AUTO_INCREMENT"`
	GameID       int64     `gorm:"not null;type:int(10)"`
	Name         string    `gorm:"not null;type:varchar(128)"`
	ChannelAppID int64     `gorm:"not null;type:int(10)"`
	LastModified time.Time `gorm:"autoUpdateTime:nano"`
}
