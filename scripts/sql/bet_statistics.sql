CREATE TABLE `bet_statistics` (
  `channel_id` int NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `game_id` int NOT NULL DEFAULT '0' COMMENT '游戏ID',
  `day` bigint NOT NULL DEFAULT '0' COMMENT '日期(当日时间戳)',
  `times` bigint DEFAULT '0' COMMENT '当日总次数',
  `bet` bigint DEFAULT '0' COMMENT '当日总下注',
  `reward` bigint DEFAULT '0' COMMENT '当日总返奖',
  `fee` bigint DEFAULT '0' COMMENT '当日总抽水',
  `win` bigint DEFAULT '0' COMMENT '当日总输赢',
  `chip` bigint DEFAULT '0' COMMENT '当日打码量',
  PRIMARY KEY (`player_id`,`game_id`,`day`) USING BTREE,
  KEY `IDX_bet_statistics_channel_id` (`channel_id`) USING BTREE,
  KEY `IDX_bet_statistics_game_id` (`game_id`),
  KEY `IDX_bet_statistics_player_id` (`day`),
  KEY `IDX_bet_statistics_day` (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家下注统计表';
