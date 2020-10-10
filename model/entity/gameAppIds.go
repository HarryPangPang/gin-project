package entity

import "time"

// GameAppIds 游戏和appId对应表
type GameAppIds struct {
	ID           uint      `gorm:"primaryKey;AUTO_INCREMENT"`
	GameID       int64     `gorm:"not null;type:int(10)"`
	Name         string    `gorm:"not null;type:varchar(128)"`
	AppID        int64     `gorm:"not null;type:int(10)"`
	LastModified time.Time `gorm:"autoUpdateTime:nano"`
}
