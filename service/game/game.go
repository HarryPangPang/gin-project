package service

func QueryAllGames() string {
	return "test"
	// return model.DB.Raw("SELECT `game`.`id`, `game`.`gameName`, `game`.`alias`, `game`.`gmId`, `game`.`bbxRegion`, `game`.`url`, `game`.`imUrl`, `game`.`accessType`, `game`.`apiSecret`, `game`.`isAutoGetData`, `game`.`sdk`, `game`.`unisdk`, `game`.`lastModified`, `gameAppIds`.`id` AS `gameAppIds.id`, `gameAppIds`.`gameId` AS `gameAppIds.gameId`, `gameAppIds`.`name` AS `gameAppIds.name`, `gameAppIds`.`appId` AS `gameAppIds.appId`, `gameAppIds`.`lastModified` AS `gameAppIds.lastModified`, `channelAppIds`.`id` AS `channelAppIds.id`, `channelAppIds`.`gameId` AS `channelAppIds.gameId`, `channelAppIds`.`name` AS `channelAppIds.name`, `channelAppIds`.`channelAppId` AS `channelAppIds.channelAppId`, `channelAppIds`.`lastModified` AS `channelAppIds.lastModified` FROM `game` AS `game` LEFT OUTER JOIN `gameAppIds` AS `gameAppIds` ON `game`.`id` = `gameAppIds`.`gameId` LEFT OUTER JOIN `channelAppIds` AS `channelAppIds` ON `game`.`id` = `channelAppIds`.`gameId` WHERE `game`.`id` = '29'")
}
