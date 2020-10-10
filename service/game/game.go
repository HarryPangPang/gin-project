package servicegame

import (
	"gmt-go/helper"
	"gmt-go/model"
	"gmt-go/model/entitys"

	"github.com/sirupsen/logrus"
)

// GameChannelAppID 当前游戏信息查询集合
type GameChannelAppID struct {
	entitys.ChannelAppIds `xorm:"extends"`
	entitys.GameAppIds    `xorm:"extends"`
	entitys.Game          `xorm:"extends"`
}

var (
	logger *logrus.Logger
)

func init() {
	logger = helper.Logger()
}

// QueryAllGames 查询所有游戏
func QueryAllGames() ([]map[string]string, error) {
	sql := "SELECT `game`.`id`, `game`.`gameName`, `game`.`alias`, `game`.`gmId`, `game`.`bbxRegion`, `game`.`url`, `game`.`imUrl`, `game`.`accessType`, `game`.`apiSecret`, `game`.`isAutoGetData`, `game`.`sdk`, `game`.`unisdk`, `game`.`lastModified`, `gameAppIds`.`id` AS `gameAppIds.id`, `gameAppIds`.`gameId` AS `gameAppIds.gameId`, `gameAppIds`.`name` AS `gameAppIds.name`, `gameAppIds`.`appId` AS `gameAppIds.appId`, `gameAppIds`.`lastModified` AS `gameAppIds.lastModified`, `channelAppIds`.`id` AS `channelAppIds.id`, `channelAppIds`.`gameId` AS `channelAppIds.gameId`, `channelAppIds`.`name` AS `channelAppIds.name`, `channelAppIds`.`channelAppId` AS `channelAppIds.channelAppId`, `channelAppIds`.`lastModified` AS `channelAppIds.lastModified` FROM `game` AS `game` LEFT OUTER JOIN `gameAppIds` AS `gameAppIds` ON `game`.`id` = `gameAppIds`.`gameId` LEFT OUTER JOIN `channelAppIds` AS `channelAppIds` ON `game`.`id` = `channelAppIds`.`gameId` WHERE `game`.`id`=?"
	results, err := model.DBM.SQL(sql, "29").QueryString()
	// games := make([]GameChannelAppID, 0)
	// logger.Errorln("查询游戏错误", 121)
	// err := model.DBM.Table("game").
	// 	Join("LEFT OUTER", "gameAppIds", "game.id = gameAppIds.gameId").
	// 	Join("LEFT OUTER", "channelAppIds", "game.id = channelAppIds.gameId").
	// 	Where("game.id=29").
	// 	Find(&games)
	if err != nil {
		logger.Errorln("查询游戏错误", err)
		return nil, err
	}
	return results, nil
}
