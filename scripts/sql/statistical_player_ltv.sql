CREATE TABLE `statistical_player_ltv` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `channel_id` int NOT NULL DEFAULT '0' COMMENT '渠道',
  `date` int NOT NULL DEFAULT '0' COMMENT '日期总额',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '枚举类型',
  `register` int NOT NULL DEFAULT '0' COMMENT '新增注册',
  `recharge` bigint NOT NULL DEFAULT '0' COMMENT '累计充值',
  `withdraw` bigint NOT NULL DEFAULT '0' COMMENT '累计提现',
  `day_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '日期内容',
  `day_content_withdraw` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '提现内容',
  `register_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '注册人员列表',
  `cost_usd` int DEFAULT '0' COMMENT '冗余cost_data.cost_usd的数据,便于/ROI的筛选查询',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `un_channel_date_range` (`channel_id`,`date`,`type`) USING BTREE,
  KEY `channel_idx` (`channel_id`) USING BTREE,
  KEY `date_idx` (`date`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1356 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户LTV';
