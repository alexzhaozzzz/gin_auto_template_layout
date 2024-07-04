CREATE TABLE `t_player_bank` (
  `n_id` bigint NOT NULL AUTO_INCREMENT,
  `n_playerID` bigint unsigned NOT NULL COMMENT '玩家ID',
  `n_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '姓名',
  `n_bankcardid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '银行卡号或UPI或IMPS账号',
  `n_bankname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '银行名称',
  `n_bankifsc` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '银行ifsc',
  `n_phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号',
  `n_createtime` datetime DEFAULT NULL COMMENT '创建时间',
  `n_isselect` tinyint DEFAULT NULL COMMENT '是否选中，默认银行卡： 1=默认,0=非默认',
  `n_razcontactid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'raz联系人ID',
  `n_email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邮箱',
  `n_panno` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `n_adno` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `n_kind` tinyint DEFAULT NULL COMMENT '银行卡类型: 0=银行卡, 1=UPI,2=IMPS',
  `n_update_time` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`n_id`) USING BTREE,
  UNIQUE KEY `t_player_bank_un` (`n_playerID`,`n_bankcardid`) USING BTREE,
  KEY `t_player_bank_n_playerID_IDX` (`n_playerID`) USING BTREE,
  KEY `t_player_bank_n_bankcardid_IDX` (`n_bankcardid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=473 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='用户银行卡定义';