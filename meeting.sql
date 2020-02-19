CREATE TABLE `meeting_meets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '活动类型 1-普通活动 2-定期活动',
  `banner` varchar(50) COMMENT '头图',
  `title` varchar(50) NOT NULL COMMENT '活动名称',
  `start_at` timestamp NOT NULL COMMENT '活动时间',
  `place` varchar(100) NOT NULL COMMENT '活动地点',
  `fee` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '活动费用',
  `person` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '计划人数',
  `content` text NULL COMMENT '活动内容',
  `state` tinyint(4) unsigned NOT NULL DEFAULT 1 COMMENT '状态 1-有效',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动';

LOCK TABLES `meeting_meets` WRITE;
INSERT INTO `meeting_meets` (`id`, `type`, `banner`, `title`, `start_at`, `place`, `fee`, `person`, `content`, `state`, `created_at`, `updated_at`, `deleted_at`)
VALUES 
    (NULL,1,NULL,'测试活动','2019-09-10 00:00:00','杭州',50,100,'这是一个测试活动','1',NULL,NULL,NULL);
UNLOCK TABLES;
