-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

--
-- テーブルのデータのダンプ `privileges`
--

INSERT INTO `privileges` (`id`, `code`, `name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'ADMIN', '管理者', 'すべての権限を持つ', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'MODERATOR', 'モデレータ', 'データの修正ができる', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'VIEWER', '閲覧者', 'データを見るだけ', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--
-- テーブルのデータのダンプ `users`
--

INSERT INTO `users` (`id`, `screen_name`,`password`, `privilege_id`, `created_at`, `updated_at`) VALUES
(1, 'tinykitten', '$2a$10$ZtDxZc1Ad418ox5FUtmu2uhQNW0hiwew5ktSmG6C1U27lkLjC7Wem', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
