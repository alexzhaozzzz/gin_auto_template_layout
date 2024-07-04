CREATE TABLE `google_analytics` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `channel_id` int NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `date` int NOT NULL DEFAULT '0' COMMENT '日期',
  `pv` int NOT NULL DEFAULT '0' COMMENT 'pv',
  `uv` int NOT NULL DEFAULT '0' COMMENT 'uv',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `un_idx` (`channel_id`,`date`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='google_analytics数据';
