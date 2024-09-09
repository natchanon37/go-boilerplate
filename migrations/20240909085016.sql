-- Create "customers" table
CREATE TABLE `customers` (
  `customer_id` varchar(64) NOT NULL,
  `name` varchar(255) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`customer_id`),
  INDEX `idx_customers_deleted_at` (`deleted_at`),
  INDEX `idx_customers_updated_at` (`updated_at` DESC)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "transactions" table
CREATE TABLE `transactions` (
  `customer_id` varchar(64) NOT NULL,
  `transaction_id` varchar(64) NOT NULL,
  `amount` varchar(64) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`customer_id`, `transaction_id`),
  INDEX `idx_transactions_deleted_at` (`deleted_at`),
  INDEX `idx_transactions_updated_at` (`updated_at` DESC),
  CONSTRAINT `fk_customers_transactions` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
