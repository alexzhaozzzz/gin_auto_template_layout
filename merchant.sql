-- MySQL dump 10.13  Distrib 8.0.37, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: merchant
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
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (44,'p','10','Wonderland','GET','','',''),(45,'p','10','Wonderland','POST','','',''),(31,'p','11','Wonderland','GET','','',''),(32,'p','11','Wonderland','POST','','',''),(69,'p','12','Wonderland','GET','','',''),(70,'p','12','Wonderland','POST','','',''),(67,'p','23','Wonderland','GET','','',''),(68,'p','24','Wonderland','GET','','',''),(71,'p','25','Wonderland','GET','','',''),(72,'p','26','Wonderland','GET','','',''),(73,'p','27','Wonderland','GET','','',''),(74,'p','29','Wonderland','GET','','',''),(75,'p','30','Wonderland','GET','','',''),(76,'p','31','Wonderland','GET','','',''),(77,'p','32','Wonderland','GET','','','');
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
  `status` smallint NOT NULL DEFAULT '0',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户渠道映射';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_channel_bind`
--

LOCK TABLES `merchant_channel_bind` WRITE;
/*!40000 ALTER TABLE `merchant_channel_bind` DISABLE KEYS */;
INSERT INTO `merchant_channel_bind` VALUES (1,10000,0,1,0,0);
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
  `guid` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(200) NOT NULL COMMENT ' 商户名称',
  `host` varchar(200) NOT NULL DEFAULT '' COMMENT '站点域名',
  `manage_host` varchar(100) NOT NULL DEFAULT '',
  `status` smallint NOT NULL DEFAULT '0',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10016 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_list`
--

LOCK TABLES `merchant_list` WRITE;
/*!40000 ALTER TABLE `merchant_list` DISABLE KEYS */;
INSERT INTO `merchant_list` VALUES (10000,'98591e4f-7afd-4284-84aa-c7e78e331734','测试商户','localhost','localhost:3000',1,0,0),(10002,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0),(10003,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0),(10004,'dadadad-dweqwe-asdasda','30qweqe','http://11111.com','http://11111.com',1,0,0),(10005,'dadadad-dweqwe-asdasda','30qweqe','http://11111.com','http://11111.com',1,0,0),(10007,'dadadad-dweqwe-asdasda','11111','http://11111.com','http://11111.com',0,0,0),(10008,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0),(10009,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0),(10010,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',0,0,0),(10011,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',0,0,0),(10012,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',0,0,0),(10013,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',0,0,0),(10014,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0),(10015,'dadadad-dweqwe-asdasda','test','http://11111.com','http://11111.com',1,0,0);
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
  `icon` varchar(50) DEFAULT '' COMMENT '图标',
  `uri` varchar(50) DEFAULT NULL COMMENT '路径',
  `show` tinyint NOT NULL DEFAULT '1' COMMENT '是否展示:1是,0否',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `permissions_ids` varchar(128) NOT NULL DEFAULT '' COMMENT '菜单绑定的权限，权限表对应的ids json数组',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (10,0,'系统','system','',1,0,'[1,2,3]',0,0),(11,10,'角色','user','/system/role',0,0,'[1,2,3]',0,0),(13,10,'用户','user','/system/user',0,0,'[1,2,3]',0,0),(14,10,'菜单','database','/system/menus',0,0,'[0]',0,0),(27,10,'权限','user','/system/permission',0,0,'[1,2,3,4,5,6]',0,0),(28,0,'仪表盘','dashBoard','/dashboard',0,0,'[1,3,4,5,2,6]',0,0),(29,0,'数据分析','report','',0,0,'[1,2,3,4,5,6]',0,0),(30,29,'综合简报','list','/report/briefReport',0,0,'[1,2,3,4,5,6]',0,0),(31,29,'渠道分析','list','/report/channel',0,0,'[1,2,3,4,5,6]',0,0),(32,0,'玩家管理','user','',0,0,'[1,3,2,4,5,6]',0,0),(33,32,'玩家列表','list','/player/player',0,0,'[1,2,3,4,5,6]',0,0),(34,32,'金币变化','list','/player/goldChange',0,0,'[1,2,3,4,5,6]',0,0),(35,32,'登录日志','list','/player/loginLog',0,0,'[1,2,3,4,5,6]',0,0),(36,32,'玩家设置','list','/player/set',0,0,'[1,2,3,4,5,6]',0,0);
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_role_menu_pk` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=256 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES (38,11,10,0,0),(39,11,11,0,0),(40,11,13,0,0),(50,13,10,0,0),(51,13,11,0,0),(52,13,13,0,0),(53,13,14,0,0),(54,13,27,0,0),(60,15,10,0,0),(61,15,11,0,0),(62,15,13,0,0),(63,15,14,0,0),(64,15,27,0,0),(65,16,10,0,0),(66,16,11,0,0),(67,16,13,0,0),(68,16,14,0,0),(69,16,27,0,0),(75,18,10,0,0),(76,18,11,0,0),(77,18,13,0,0),(78,18,14,0,0),(79,18,27,0,0),(80,19,10,0,0),(81,19,11,0,0),(82,19,13,0,0),(83,19,14,0,0),(84,19,27,0,0),(85,20,10,0,0),(86,20,11,0,0),(87,20,13,0,0),(88,20,14,0,0),(89,20,27,0,0),(90,21,10,0,0),(91,21,11,0,0),(92,21,13,0,0),(93,21,14,0,0),(94,21,27,0,0),(95,22,10,0,0),(96,22,11,0,0),(97,22,13,0,0),(98,22,14,0,0),(99,22,27,0,0),(109,17,10,0,0),(110,17,11,0,0),(111,17,13,0,0),(112,17,14,0,0),(113,17,27,0,0),(114,17,28,0,0),(138,10,1,0,0),(139,10,2,0,0),(140,10,3,0,0),(181,23,28,0,0),(182,23,29,0,0),(183,23,30,0,0),(184,23,31,0,0),(185,23,32,0,0),(186,23,33,0,0),(187,23,34,0,0),(188,23,35,0,0),(189,23,36,0,0),(190,24,28,0,0),(191,24,29,0,0),(192,24,30,0,0),(193,24,31,0,0),(194,24,32,0,0),(195,24,33,0,0),(196,24,34,0,0),(197,24,35,0,0),(198,24,36,0,0),(204,12,0,0,0),(205,12,1,0,0),(206,12,2,0,0),(207,12,3,0,0),(208,26,0,0,0),(209,26,1,0,0),(210,26,2,0,0),(211,26,3,0,0),(212,26,4,0,0),(213,26,5,0,0),(214,26,6,0,0),(215,26,7,0,0),(216,26,8,0,0),(217,26,9,0,0),(218,26,10,0,0),(219,26,11,0,0),(220,26,12,0,0),(221,26,13,0,0),(222,27,0,0,0),(223,27,1,0,0),(224,27,2,0,0),(225,27,3,0,0),(226,27,4,0,0),(227,27,5,0,0),(228,27,6,0,0),(229,27,7,0,0),(230,27,8,0,0),(231,27,9,0,0),(232,27,10,0,0),(233,27,11,0,0),(234,27,12,0,0),(235,27,13,0,0),(236,31,0,0,0),(237,31,1,0,0),(238,31,2,0,0),(239,31,3,0,0),(240,31,4,0,0),(241,31,5,0,0),(242,32,10,0,0),(243,32,11,0,0),(244,32,13,0,0),(245,32,14,0,0),(246,32,27,0,0),(247,32,28,0,0),(248,32,29,0,0),(249,32,30,0,0),(250,32,31,0,0),(251,14,0,0,0),(252,14,1,0,0),(253,14,2,0,0),(254,14,3,0,0),(255,14,4,0,0);
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_role_permissions_pk` (`role_id`,`permission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色权限';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_permissions`
--

LOCK TABLES `sys_role_permissions` WRITE;
/*!40000 ALTER TABLE `sys_role_permissions` DISABLE KEYS */;
INSERT INTO `sys_role_permissions` VALUES (30,11,1,0,0),(31,11,2,0,0),(32,11,3,0,0),(33,11,4,0,0),(43,10,1,0,0),(44,10,2,0,0),(45,10,3,0,0),(66,23,1,0,0),(67,24,1,0,0),(68,12,1,0,0),(69,12,3,0,0),(70,25,1,0,0),(71,26,1,0,0),(72,27,1,0,0),(73,29,1,0,0),(74,30,1,0,0),(75,31,1,0,0),(76,32,1,0,0);
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
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_role_users_pk` (`role_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_users`
--

LOCK TABLES `sys_role_users` WRITE;
/*!40000 ALTER TABLE `sys_role_users` DISABLE KEYS */;
INSERT INTO `sys_role_users` VALUES (7,3,0,0,0),(15,5,1,0,0),(17,6,6,0,0),(18,10,7,0,0),(19,10,8,0,0),(20,10,9,0,0),(21,10,10,0,0),(22,10,11,0,0),(23,10,12,0,0),(24,10,13,0,0),(25,10,14,0,0),(26,10,15,0,0),(27,10,16,0,0),(28,10,17,0,0),(29,10,18,0,0);
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
  `guid` varchar(100) NOT NULL DEFAULT '',
  `merchant_id` bigint NOT NULL DEFAULT '0' COMMENT '商户ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色标识',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_roles`
--

LOCK TABLES `sys_roles` WRITE;
/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT INTO `sys_roles` VALUES (10,'bbbbb',10000,'超级管理员12','super12',0,0),(11,'ccccc',10000,'测试1','super',0,0),(12,'aa',10000,'01','01',0,0),(13,'aa',10000,'02','02',0,0),(14,'aa',10000,'03','03',0,0),(15,'aa',10000,'04','04',0,0),(16,'aa',10000,'05','05',0,0),(17,'aa',10000,'06','06',0,0),(18,'aa',10000,'07','07',0,0),(19,'aa',10000,'08','08',0,0),(20,'aa',10000,'09','09',0,0),(21,'aa',10000,'10','10',0,0),(22,'aa',10000,'11','11',0,0),(23,'aa',10000,'01','01',0,0),(24,'aa',10000,'01','01',0,0),(25,'aa',10000,'888','888',0,0),(26,'aa',10000,'888','888',0,0),(27,'aa',10000,'888','888',0,0),(28,'aa',10000,'888','888',0,0),(29,'aa',10000,'999','999',0,0),(30,'aa',10000,'333','333',0,0),(31,'aa',10000,'333','333',0,0),(32,'aa',10000,'444','444',0,0);
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
  `guid` varchar(100) NOT NULL DEFAULT '',
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick_name` varchar(100) NOT NULL DEFAULT '',
  `merchant_id` bigint NOT NULL DEFAULT '0',
  `avatar` varchar(500) DEFAULT NULL,
  `email` varchar(100) NOT NULL DEFAULT '',
  `phone` varchar(100) NOT NULL DEFAULT '',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  `last_login_time` int NOT NULL DEFAULT '0',
  `last_login_out_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_pk` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'','evan','bbb','cccccc',10000,NULL,'evan','evan',0,0,1718446605,0),(6,'','alex1','1111','alex',10000,'','admin@123.com','alex',0,0,0,0),(7,'','11','698d51a19d8a121ce581499d7b701668','11',10000,'','admin@123.com','18922222222',0,0,0,0),(8,'','12','698d51a19d8a121ce581499d7b701668','12',10000,'','admin@123.com','18922222222',0,0,0,0),(9,'','13','698d51a19d8a121ce581499d7b701668','13',10000,'','admin@123.com','18922222222',0,0,0,0),(10,'','14','698d51a19d8a121ce581499d7b701668','14',10000,'','admin@123.com','18922222222',0,0,0,0),(11,'','15','698d51a19d8a121ce581499d7b701668','15',10000,'','admin@123.com','18922222222',0,0,0,0),(12,'','16','698d51a19d8a121ce581499d7b701668','16',10000,'','admin@123.com','18922222222',0,0,0,0),(13,'','17','698d51a19d8a121ce581499d7b701668','17',10000,'','admin@123.com','18922222222',0,0,0,0),(14,'','18','698d51a19d8a121ce581499d7b701668','18',10000,'','admin@123.com','18922222222',0,0,0,0),(15,'','19','698d51a19d8a121ce581499d7b701668','19',10000,'','admin@123.com','18922222222',0,0,0,0),(16,'','20','698d51a19d8a121ce581499d7b701668','20',10000,'','admin@123.com','18922222222',0,0,0,0),(17,'','21','698d51a19d8a121ce581499d7b701668','21',10000,'','admin@123.com','18922222222',0,0,0,0),(18,'','22','698d51a19d8a121ce581499d7b701668','22',10000,'','admin@123.com','18922222222',0,0,0,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-27  8:56:27
