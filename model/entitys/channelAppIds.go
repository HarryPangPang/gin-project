package entitys

import "time"

// ChannelAppIds 游戏和channelappId对应表
type ChannelAppIds struct {
	ID           int64     `xorm:"not null INT(10) pk autoincr"`
	GameID       int       `xorm:"not null UNSIGNED INT(10)"`
	Name         string    `xorm:"not null VARCHAR(128)"`
	ChannelAppID int       `xorm:"not null INT(10)"`
	LastModified time.Time `xorm:"updated not null"`
}
