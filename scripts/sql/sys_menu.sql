

CREATE TABLE `sys_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `uri` varchar(50) DEFAULT NULL COMMENT '路径',
  `show` tinyint NOT NULL DEFAULT '1' COMMENT '是否展示:1是,0否',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单列表';

