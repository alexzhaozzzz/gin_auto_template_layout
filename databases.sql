-- MySQL dump 10.13  Distrib 8.0.37, for Linux (x86_64)
--
-- Host: 192.168.1.22    Database: merchant
-- ------------------------------------------------------
-- Server version	8.0.37-0ubuntu0.22.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (16,'p','10','Wonderland','GET','','',''),(17,'p','10','Wonderland','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merchant_channel_bind`
--

DROP TABLE IF EXISTS `merchant_channel_bind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `merchant_channel_bind` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchant_id` int NOT NULL COMMENT '商户 Id',
  `channel_id` int NOT NULL COMMENT '渠道 ID',
  `status` smallint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户渠道映射';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_channel_bind`
--

LOCK TABLES `merchant_channel_bind` WRITE;
/*!40000 ALTER TABLE `merchant_channel_bind` DISABLE KEYS */;
INSERT INTO `merchant_channel_bind` VALUES (1,10000,123123,1,0,0);
/*!40000 ALTER TABLE `merchant_channel_bind` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merchant_list`
--

DROP TABLE IF EXISTS `merchant_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_list`
--

LOCK TABLES `merchant_list` WRITE;
/*!40000 ALTER TABLE `merchant_list` DISABLE KEYS */;
INSERT INTO `merchant_list` VALUES (10000,'98591e4f-7afd-4284-84aa-c7e78e331734','测试商户','localhost','localhost:3000',1,0,0);
/*!40000 ALTER TABLE `merchant_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `uri` varchar(50) DEFAULT NULL COMMENT '路径',
  `show` tinyint NOT NULL DEFAULT '1' COMMENT '是否展示:1是,0否',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `permissions_ids` varchar(128) NOT NULL DEFAULT '' COMMENT '菜单绑定的权限，权限表对应的ids json数组',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (10,0,'系统','system','',1,0,'[1,2,3]',0,0),(11,10,'角色','user','/system/role',0,0,'[1,2,3]',0,0),(13,10,'用户','user','/system/user',0,0,'[1,2,3]',0,0),(14,10,'菜单','menu','/system/menus',0,0,'[1,2,3]',0,0),(25,0,'test11','user','/system/user',0,0,'[1]',0,0),(26,0,'test2','user','/system/user',0,0,'[1,2]',0,0);
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menu`
--

DROP TABLE IF EXISTS `sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES (29,10,1,0,0),(30,10,2,0,0),(31,10,3,0,0);
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_permissions`
--

DROP TABLE IF EXISTS `sys_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_permissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` bigint NOT NULL,
  `permission_id` bigint NOT NULL,
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色权限';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_permissions`
--

LOCK TABLES `sys_role_permissions` WRITE;
/*!40000 ALTER TABLE `sys_role_permissions` DISABLE KEYS */;
INSERT INTO `sys_role_permissions` VALUES (17,10,1,0,0),(18,10,2,0,0),(19,10,3,0,0);
/*!40000 ALTER TABLE `sys_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_users`
--

DROP TABLE IF EXISTS `sys_role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_users`
--

LOCK TABLES `sys_role_users` WRITE;
/*!40000 ALTER TABLE `sys_role_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_roles`
--

DROP TABLE IF EXISTS `sys_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `guid` varchar(100) NOT NULL,
  `merchant_id` bigint NOT NULL COMMENT '商户ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色标识',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_roles`
--

LOCK TABLES `sys_roles` WRITE;
/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT INTO `sys_roles` VALUES (6,'aa',10000,'超级管理员','super',0,0),(10,'aa',10000,'测试1','test_1',0,0);
/*!40000 ALTER TABLE `sys_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `guid` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick_name` varchar(100) NOT NULL,
  `merchant_id` bigint NOT NULL,
  `role` varchar(200) NOT NULL,
  `avatar` varchar(500) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `is_master` smallint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  `last_login_time` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'','evan','e10adc3949ba59abbe56e057f20f883e','evan',10000,'1',NULL,'','',1,0,0,1718446605);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_test`
--

DROP TABLE IF EXISTS `user_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_test` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `guid` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick_name` varchar(100) NOT NULL,
  `merchant_id` bigint NOT NULL,
  `role` varchar(200) NOT NULL,
  `avatar` varchar(500) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `is_master` smallint NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  `last_login_time` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_test`
--

LOCK TABLES `user_test` WRITE;
/*!40000 ALTER TABLE `user_test` DISABLE KEYS */;
INSERT INTO `user_test` VALUES (1,'','evan','e10adc3949ba59abbe56e057f20f883e','evan',10000,'1',NULL,'','',1,0,0,1718446605);
/*!40000 ALTER TABLE `user_test` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-24 12:32:26
