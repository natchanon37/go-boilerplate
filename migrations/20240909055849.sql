-- Create "customers" table
CREATE TABLE `customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_customers_deleted_at` (`deleted_at`),
  INDEX `idx_customers_updated_at` (`updated_at` DESC)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "transactions" table
CREATE TABLE `transactions` (
  `transaction_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `customer_id` bigint unsigned NULL,
  `amount` decimal(10,2) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`transaction_id`),
  INDEX `fk_customers_transactions` (`customer_id`),
  INDEX `idx_transactions_deleted_at` (`deleted_at`),
  INDEX `idx_transactions_updated_at` (`updated_at` DESC),
  CONSTRAINT `fk_customers_transactions` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
