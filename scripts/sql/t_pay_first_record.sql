CREATE TABLE `t_pay_first_record` (
  `n_id` bigint NOT NULL AUTO_INCREMENT,
  `n_uid` bigint DEFAULT NULL COMMENT '玩家ID',
  `n_channel` int DEFAULT NULL COMMENT '渠道ID',
  `n_coin` bigint DEFAULT NULL COMMENT '首次充值前的金币',
  `n_amount` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '首次充值额',
  `n_order` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '首次充值订单ID',
  `n_create_time` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`n_id`) USING BTREE,
  UNIQUE KEY `t_pay_first_record_un` (`n_uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=650 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='首次充值记录表';
