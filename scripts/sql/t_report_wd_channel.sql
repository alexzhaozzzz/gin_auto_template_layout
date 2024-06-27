CREATE TABLE `t_report_wd_channel` (
  `n_id` bigint NOT NULL AUTO_INCREMENT,
  `n_pid` int DEFAULT NULL COMMENT '通道ID',
  `n_pname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '通道名称',
  `n_submit` bigint DEFAULT NULL COMMENT '请求次数',
  `n_success` bigint DEFAULT NULL COMMENT '成功次数',
  `n_fail` bigint DEFAULT NULL COMMENT '失败次数',
  `n_err` bigint DEFAULT NULL COMMENT '错误次数',
  `n_create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `n_point` int DEFAULT NULL COMMENT '点',
  PRIMARY KEY (`n_id`) USING BTREE,
  KEY `t_report_pay_channel_n_pid_IDX` (`n_pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62986 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='充值通道统计信息';
