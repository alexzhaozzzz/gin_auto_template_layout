CREATE TABLE `statistical_rank_pay` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `channel_id` bigint DEFAULT NULL COMMENT '渠道ID',
  `player_id` bigint DEFAULT NULL COMMENT '玩家ID',
  `time_int` int DEFAULT NULL COMMENT '当天时间缀',
  `pay` bigint DEFAULT NULL COMMENT '当天累计充值',
  `withdraw` bigint DEFAULT NULL COMMENT '当日累计提现',
  `pay_number` int DEFAULT NULL COMMENT '充值笔数',
  `gold` bigint DEFAULT NULL COMMENT '真金+绑金',
  `currency_recharge` bigint DEFAULT NULL COMMENT '充值总额_当天',
  `currency_withdraw` bigint DEFAULT NULL COMMENT '总提现额_当天',
  `n_src` tinyint DEFAULT NULL COMMENT '充值来源（最后一笔充值',
  `n_last_theme` smallint DEFAULT NULL COMMENT 'spin主题（最后一笔充值前spin主题）',
  `beginTime` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '最后登录',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='每日充值榜';
