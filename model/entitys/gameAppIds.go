package entitys

import "time"

// GameAppIds 游戏和appId对应表
type GameAppIds struct {
	ID           int64     `xorm:"not null INT(10) pk autoincr"`
	GameID       int       `xorm:"not null UNSIGNED INT(10)"`
	Name         string    `xorm:"not null VARCHAR(128)"`
	AppID        int       `xorm:"not null INT(10)"`
	LastModified time.Time `xorm:"updated not null"`
}
