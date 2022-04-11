CREATE TABLE `customers` (
  `id` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT current_timestamp()
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;


CREATE TABLE `merchants` (
  `id` varchar(255) NOT NULL,
  `merchant_name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT current_timestamp()
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;


CREATE TABLE `products` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `detail` text DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `stock` int(11) DEFAULT NULL,
  `merchant_id` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT current_timestamp()
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;


CREATE TABLE `roles` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT 'user',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;


INSERT INTO `roles` (`id`, `name`, `created_at`, `updated_at`) VALUES
(0, 'customer', '2022-03-30 10:50:59', '2022-03-30 10:50:59'),
(1, 'merchant', '2022-03-30 10:51:19', '2022-03-30 10:51:19');


CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;


CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `hash_password` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB;



ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `role_id` (`role_id`);


ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `role_id` (`role_id`);


ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `merchant_id` (`merchant_id`);


ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);


ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `role_id` (`role_id`);

ALTER TABLE `roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;


ALTER TABLE `customers`
  ADD CONSTRAINT `customers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `customers_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);


ALTER TABLE `merchants`
  ADD CONSTRAINT `merchants_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `merchants_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);


ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`id`);


ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);