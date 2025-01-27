CREATE TABLE `t_player` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `accountID` bigint NOT NULL COMMENT '账户ID',
  `playerID` bigint NOT NULL COMMENT '玩家ID',
  `showUID` bigint NOT NULL COMMENT '显示ID',
  `type` int NOT NULL DEFAULT '1' COMMENT '玩家类型1.普通玩家，2.机器人，3.QA\n2.\n3.',
  `channelID` int DEFAULT NULL COMMENT '渠道ID',
  `productID` int DEFAULT '9999' COMMENT '产品ID',
  `inviteID` bigint unsigned DEFAULT NULL COMMENT '邀请人玩家playerID',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` int DEFAULT '1' COMMENT '性别',
  `avatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '头像地址',
  `hdAvatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '高清头像地址',
  `provinceID` int DEFAULT NULL COMMENT '省ID',
  `cityID` int DEFAULT NULL COMMENT '市ID',
  `areaID` int DEFAULT NULL COMMENT '地区ID',
  `name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '真实姓名',
  `phone` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '手机号码',
  `phoneType` int NOT NULL DEFAULT '1' COMMENT '手机类型:1.正常手机号，2.黑号',
  `phoneLocation` int NOT NULL DEFAULT '0' COMMENT '手机号归属地',
  `idCard` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '身份证',
  `wxOpenId` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '游客账号ID, 微信opendID或googleID，第三方账号ID',
  `wxUnionId` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '微信unionID',
  `wxName` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `isWhiteList` tinyint(1) DEFAULT '0' COMMENT '是否QA，默认否',
  `zipCode` int DEFAULT NULL COMMENT '邮编',
  `shippingAddr` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '收获地址',
  `status` int DEFAULT '1' COMMENT '1可登录，2冻结，默认为1',
  `remark` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `version` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端版本号',
  `device` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端设备名称',
  `imei` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端设备唯一识别码',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间，通常也是注册时间',
  `createBy` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '更新人',
  `promoterID` varchar(70) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '推广员ID',
  `fromID` varchar(70) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '来源ID',
  `zoneID` varchar(70) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '场景ID',
  `topVersion` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端最高版本号',
  `regip` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '注册IP',
  `topInvite` bigint NOT NULL DEFAULT '0' COMMENT '最高上级玩家id',
  `money_pass` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '玩家资金密码',
  `vip_level` tinyint DEFAULT '0' COMMENT 'vip等级',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `t_player_1_UN_playerID` (`playerID`) USING BTREE,
  UNIQUE KEY `t_player_1_UN_showUID` (`showUID`) USING BTREE,
  UNIQUE KEY `t_player_1_UN_ACCOUTID` (`accountID`) USING BTREE,
  KEY `t_player_imei_IDX` (`imei`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3276827 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='玩家表';
