CREATE TABLE `statistical_log_game` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `channel_id` int NOT NULL DEFAULT '0' COMMENT '渠道',
  `game_id` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏ID',
  `date` int NOT NULL DEFAULT '0' COMMENT '日期',
  `platform` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台',
  `join_num` bigint NOT NULL DEFAULT '0' COMMENT '参与人数',
  `bet` bigint NOT NULL DEFAULT '0' COMMENT '下注',
  `reward` bigint NOT NULL DEFAULT '0' COMMENT '返奖',
  `win` bigint NOT NULL DEFAULT '0' COMMENT '输赢',
  `chip` bigint NOT NULL DEFAULT '0' COMMENT '打码量',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `un_idx` (`channel_id`,`date`,`game_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1652 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='主题统计';
