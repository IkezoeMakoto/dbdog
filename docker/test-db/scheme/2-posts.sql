CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT 'ユーザID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'タイトル',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '本文',
  `published_at` timestamp NULL DEFAULT NULL COMMENT '公開日時',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '削除日時',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `posts_fk1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投稿';
