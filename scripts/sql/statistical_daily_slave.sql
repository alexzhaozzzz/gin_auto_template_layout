CREATE TABLE `statistical_daily_slave` (
  `s_d_id` int unsigned NOT NULL COMMENT '1:1 statistical_daily (id)',
  `register_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '注册玩家id',
  `login_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '登录玩家id',
  `gold_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '活跃玩家id',
  `spin_active` int DEFAULT '0' COMMENT '新增活跃',
  `arpu` int DEFAULT '0' COMMENT 'arpu活跃帐号的充值金额',
  `new_arpu` int DEFAULT '0' COMMENT '新用户的当天充值金额',
  `old_arpu` int DEFAULT '0' COMMENT '老用户(非当天注册)当天充值金额',
  `old_arpu_recharge` int NOT NULL DEFAULT '0' COMMENT '老用户(非当天首充)当天充值金额',
  `active_withdraw` int DEFAULT '0' COMMENT '活跃提现',
  `new_withdraw` int DEFAULT '0' COMMENT '新用户提现',
  `old_withdraw` int DEFAULT '0' COMMENT '老用户提现',
  `coins` bigint DEFAULT '0' COMMENT '非绑定金币',
  `ingots` bigint DEFAULT '0' COMMENT '绑定金币',
  `ingots_pay` bigint NOT NULL DEFAULT '0' COMMENT '付费用户绑金',
  `n_bonus_sum` bigint DEFAULT '0' COMMENT '奖金',
  `new_pay_limit` bigint NOT NULL DEFAULT '0' COMMENT '新增充值额度',
  `recharge_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '付费充值的玩家ID',
  `withdraw_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '提现玩家ID',
  `first_recharge_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '首充玩家ID',
  `bankruptcy` bigint NOT NULL DEFAULT '0' COMMENT '破产玩家数量',
  `new_pay_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '新注册付费玩家',
  `re_new_pay_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '复购新注册付费玩家',
  `new_device_login_num` int NOT NULL DEFAULT '0' COMMENT '新增设备登录',
  `old_device_login_num` int NOT NULL DEFAULT '0' COMMENT '老设备登录',
  `recharge_success_rate` decimal(9,2) NOT NULL COMMENT '支付成功率',
  `new_recharge_success_rate` decimal(9,2) NOT NULL COMMENT '新用户增支付成功率',
  `old_recharge_success_rate` decimal(9,2) NOT NULL COMMENT '老用户付费成功率',
  `new_recharge_success_rate_curr` decimal(9,2) NOT NULL COMMENT '新用户增支付成功率(新增用户支付成功订单/ 新增用户所有支付订单)',
  `old_recharge_success_rate_curr` decimal(9,2) NOT NULL COMMENT '老用户付费成功率(老用户支付成功订单/老用户所有支付订单)',
  `install_src_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '安装来源',
  `not_sdk_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '非sdk真金模式用户',
  `re_old_pay_content` longtext COLLATE utf8mb4_general_ci COMMENT '复购老用户',
  PRIMARY KEY (`s_d_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
