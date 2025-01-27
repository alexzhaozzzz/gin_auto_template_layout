CREATE TABLE `t_player_device` (
  `n_id` bigint NOT NULL AUTO_INCREMENT,
  `n_player_id` bigint NOT NULL COMMENT '玩家ID',
  `n_channel` int DEFAULT NULL COMMENT '渠道ID',
  `n_reg_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '注册IP',
  `n_client_ver` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户端版本号',
  `n_os` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '系统信息',
  `n_dev_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备唯一ID',
  `n_location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '玩家所在地区',
  `n_coordinate` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '坐标',
  `n_lang` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机系统语言',
  `n_phone_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机型号',
  `n_sdk_data` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'leo sdk data',
  `n_sp` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机运营商',
  `n_is_real` tinyint DEFAULT NULL COMMENT '是否是真金用户',
  `n_adid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ad id',
  `n_reg_time` datetime DEFAULT NULL COMMENT '注册时间',
  `n_last_login_time` bigint DEFAULT '0' COMMENT '最后登录时间，unix时间戳',
  PRIMARY KEY (`n_id`) USING BTREE,
  UNIQUE KEY `t_player_context_UN` (`n_player_id`) USING BTREE,
  KEY `t_player_device_n_channel_IDX` (`n_channel`,`n_reg_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3258 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家设备信息表';
