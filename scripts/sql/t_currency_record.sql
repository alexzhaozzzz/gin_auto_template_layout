CREATE TABLE `t_currency_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tradeID` varchar(64) NOT NULL COMMENT '流水ID',
  `playerID` bigint NOT NULL COMMENT '玩家ID',
  `channel` int DEFAULT NULL COMMENT '渠道ID',
  `currencyType` int DEFAULT NULL COMMENT '货币类型: 1=金币, 2=元宝（钻石）， 3=房卡',
  `amount` bigint DEFAULT NULL COMMENT '加减值',
  `beforeBalance` bigint DEFAULT NULL COMMENT '操作前金币值',
  `afterBalance` bigint DEFAULT NULL COMMENT '操作后金币值',
  `tradeTime` datetime DEFAULT NULL COMMENT '创建时间',
  `status` tinyint(1) DEFAULT NULL COMMENT '操作结果： 1=成功，0=失败',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `gameId` bigint DEFAULT NULL COMMENT '游戏ID',
  `level` int DEFAULT NULL COMMENT '场次ID',
  `funcId` int DEFAULT NULL COMMENT '行为ID或功能ID',
  `productID` bigint DEFAULT '9999' COMMENT '产品ID。 红幺鸡=9999，小程序=10000',
  PRIMARY KEY (`id`),
  KEY `t_currency_record_playerID_IDX` (`playerID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 COMMENT='金币流水表';
