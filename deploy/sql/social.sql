CREATE DATABASE IF NOT EXISTS `social` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `social`;
CREATE TABLE `friends` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '好友关系唯一ID',
   `user_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户ID',
   `friend_uid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '好友的用户ID',
   `remark` varchar(255) DEFAULT NULL COMMENT '好友备注',
   `add_source` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '添加来源(1:搜索 2:群组 3:二维码 4:名片)',
   `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='好友关系表';

CREATE TABLE `friend_requests` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '好友请求唯一ID',
   `user_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '被请求用户ID',
   `req_uid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方用户ID',
   `req_msg` varchar(255) DEFAULT NULL COMMENT '好友请求信息',
   `req_time` timestamp NOT NULL COMMENT '请求时间',
   `handle_result` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '处理结果(0:未处理 1:同意 2:拒绝)',
   `handle_msg` varchar(255) DEFAULT NULL COMMENT '处理信息',
   `handled_at` timestamp NULL DEFAULT NULL COMMENT '处理时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='好友请求表';

CREATE TABLE `groups` (
 `id` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群组ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群名称',
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群图标',
  `status` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '群状态(0:正常 1:解散)',
  `creator_uid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建人用户ID',
  `group_type` int(11) NOT NULL COMMENT '群类型(1:普通群 2:企业群)',
  `is_verify` boolean NOT NULL COMMENT '入群验证',
  `notification` varchar(255) DEFAULT NULL COMMENT '群公告',
  `notification_uid` varchar(64) DEFAULT NULL COMMENT '公告发布人ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群组信息表';

CREATE TABLE `group_members` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '群成员关系唯一ID',
 `group_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群组ID',
 `user_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户ID',
 `role_level` tinyint COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '成员角色(1:普通成员 2:管理员 3:群主)',
 `join_time` timestamp NULL DEFAULT NULL COMMENT '加入时间',
 `join_source` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '加入来源(1:邀请 2:搜索 3:二维码)',
 `inviter_uid` varchar(64) DEFAULT NULL COMMENT '邀请人ID',
 `operator_uid` varchar(64) DEFAULT NULL COMMENT '操作人ID(用于管理员操作)',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群成员表';

CREATE TABLE `group_requests` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '加群请求唯一ID',
  `req_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求ID',
  `group_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群组ID',
  `req_msg` varchar(255) DEFAULT NULL COMMENT '请求信息',
  `req_time` timestamp NULL DEFAULT NULL COMMENT '请求时间',
  `join_source` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '加入来源(1:邀请 2:搜索 3:二维码)',
  `inviter_user_id` varchar(64) DEFAULT NULL COMMENT '邀请人ID',
  `handle_user_id` varchar(64) DEFAULT NULL COMMENT '处理人ID',
  `handle_time` timestamp NULL DEFAULT NULL COMMENT '处理时间',
  `handle_result` tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '处理结果(0:未处理 1:同意 2:拒绝)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='加群请求表';