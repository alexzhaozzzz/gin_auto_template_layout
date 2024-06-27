CREATE TABLE `merchant_channel_bind` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchant_id` int NOT NULL COMMENT '商户 Id',
  `channel_id` int NOT NULL COMMENT '渠道 ID',
  `status` smallint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户渠道映射';
