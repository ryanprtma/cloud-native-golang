CREATE TABLE `products` (
  `id` varchar(255) PRIMARY KEY,
  `name` varchar(255),
  `detail` text(255),
  `price` int(11),
  `stock` int(11),
  `merchant_id` varchar (255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT (now())
);

ALTER TABLE `products` ADD FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`id`);