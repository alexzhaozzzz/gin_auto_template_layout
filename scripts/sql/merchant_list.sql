

CREATE TABLE `merchant_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(100) NOT NULL,
  `name` varchar(200) NOT NULL COMMENT ' 商户名称',
  `host` varchar(200) NOT NULL COMMENT '站点域名',
  `manage_host` varchar(100) NOT NULL,
  `status` smallint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户表';

