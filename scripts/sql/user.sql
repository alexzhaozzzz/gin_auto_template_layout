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
