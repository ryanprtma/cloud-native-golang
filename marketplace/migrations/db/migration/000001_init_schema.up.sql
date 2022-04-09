CREATE TABLE `roles` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) DEFAULT "user",
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `users` (
  `id` varchar(255) PRIMARY KEY,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `email` varchar(255),
  `role_id` int(11),
  `hash_password` varchar(255),
  `token` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `merchants` (
  `id` varchar(255) PRIMARY KEY,
  `merchant_name` varchar(255),
  `address` varchar(255),
  `user_id` varchar(255),
  `role_id` int(11),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT (now())
);

CREATE TABLE `customers` (
  `id` varchar(255) PRIMARY KEY,
  `address` varchar(255),
  `user_id` varchar(255),
  `role_id` int(11)
);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `merchants` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `merchants` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `customers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `customers` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);