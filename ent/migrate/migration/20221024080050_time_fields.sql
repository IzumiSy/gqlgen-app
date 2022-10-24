-- add column "create_time" to table: "todos"
ALTER TABLE `todos` ADD COLUMN `create_time` datetime NOT NULL;
-- add column "update_time" to table: "todos"
ALTER TABLE `todos` ADD COLUMN `update_time` datetime NOT NULL;
-- add column "create_time" to table: "users"
ALTER TABLE `users` ADD COLUMN `create_time` datetime NOT NULL;
-- add column "update_time" to table: "users"
ALTER TABLE `users` ADD COLUMN `update_time` datetime NOT NULL;
-- add column "create_time" to table: "categories"
ALTER TABLE `categories` ADD COLUMN `create_time` datetime NOT NULL;
