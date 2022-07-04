CREATE DATABASE IF NOT EXISTS `mercado-fresh-panic`;

USE `mercado-fresh-panic`;

DROP TABLE IF EXISTS `countries`;

CREATE TABLE `countries` (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  country_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `countries` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `provinces`;

CREATE TABLE `provinces`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  province_name VARCHAR(255) NOT NULL,
  id_country_fk BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (id_country_fk) REFERENCES countries(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `provinces` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `localities`;

CREATE TABLE `localities` (
  id VARCHAR(255) NOT NULL,
  locality_name VARCHAR(255) NOT NULL,
  province_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (province_id) REFERENCES provinces(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `provinces` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `sellers`;

CREATE TABLE `sellers` (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cid BIGINT (64) UNIQUE NOT NULL,
  company_name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  telephone VARCHAR(255) NOT NULL,
  locality_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (locality_id) REFERENCES localities(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `sellers` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `carriers`;

CREATE TABLE `carriers` (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cid VARCHAR(255) UNIQUE NOT NULL,
  company_name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  telephone VARCHAR(255) NOT NULL,
  locality_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (locality_id) REFERENCES localities(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `carriers` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `warehouses`;

CREATE TABLE `warehouses` (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  address VARCHAR(255) NOT NULL,
  telephone VARCHAR(255) NOT NULL,
  warehouse_code VARCHAR(255) UNIQUE NOT NULL,
  locality_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (locality_id) REFERENCES localities(id),
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `warehouses` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `products_types`;

CREATE TABLE `products_types` (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `products_types` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NOT NULL,
  expiration_rate DECIMAL(19, 2) NOT NULL,
  freezing_rate DECIMAL(19, 2) NOT NULL,
  height DECIMAL(19, 2) NOT NULL,
  length DECIMAL(19, 2) NOT NULL,
  net_weight DECIMAL(19, 2) NOT NULL,
  product_code VARCHAR(255) NOT NULL,
  recommended_freezing_temperature DECIMAL(19, 2) NOT NULL,
  width DECIMAL(19, 2) NOT NULL,
  product_type BIGINT UNSIGNED NOT NULL,
  seller_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (product_type) REFERENCES products_types(id),
  FOREIGN KEY (seller_id) REFERENCES sellers(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `products` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `sections`;

CREATE TABLE `sections`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  section_number BIGINT NOT NULL,
  current_capacity BIGINT UNSIGNED NOT NULL,
  current_temperature DECIMAL(19, 2) NOT NULL,
  maximum_capacity BIGINT UNSIGNED NOT NULL,
  minimum_capacity BIGINT UNSIGNED NOT NULL,
  minimum_temperature DECIMAL(19, 2) NOT NULL,
  product_type BIGINT UNSIGNED NOT NULL,
  warehouse_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (product_type) REFERENCES products_types(id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `sections` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `product_batches`;

CREATE TABLE `product_batches`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  batch_number BIGINT NOT NULL,
  current_quantity BIGINT UNSIGNED NOT NULL,
  current_temperature DECIMAL(19, 2) NOT NULL,
  due_date DATETIME(6) NOT NULL,
  initial_quantity BIGINT UNSIGNED NOT NULL,
  manufacturing_date DATETIME(6) NOT NULL,
  manufacturing_hour DATETIME(6) NOT NULL,
  minimum_temperature DECIMAL(19, 2) NOT NULL,
  product_id BIGINT UNSIGNED NOT NULL,
  section_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (section_id) REFERENCES sections(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `product_batches` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `product_records`;

CREATE TABLE `product_records`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  last_update_date DATETIME(6) NOT NULL,
  purchase_price DECIMAL(19, 2) NOT NULL,
  sale_price DECIMAL(19, 2) NOT NULL,
  product_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `product_records` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `employees`;

CREATE TABLE `employees`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  id_card_number VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  warehouse_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `employees` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `inbound_orders`;

CREATE TABLE `inbound_orders`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  order_date DATETIME(6) NOT NULL,
  order_number VARCHAR(255) NOT NULL,
  employee_id BIGINT UNSIGNED NOT NULL,
  product_batch_id BIGINT UNSIGNED NOT NULL,
  warehouse_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (employee_id) REFERENCES employees(id),
  FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLE `inbound_orders` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `order_status`;

CREATE TABLE `order_status`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `order_status` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `buyers`;

CREATE TABLE `buyers`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  id_card_number VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `buyers` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `purchase_orders`;

CREATE TABLE `purchase_orders`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  order_number VARCHAR(255) NOT NULL,
  order_date DATETIME(6) NOT NULL,
  tracking_code VARCHAR(255) NOT NULL,
  buyer_id BIGINT UNSIGNED NOT NULL,
  order_status_id BIGINT UNSIGNED NOT NULL,
  product_record_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY (buyer_id) REFERENCES buyers(id),
  FOREIGN KEY (order_status_id) REFERENCES order_status(id),
  FOREIGN KEY (product_record_id) REFERENCES product_records(id),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `purchase_orders` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  password VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLE `users` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `rol`;

CREATE TABLE `rol`(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NOT NULL,
  rol_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `rol` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `user_rol`;

CREATE TABLE `user_rol`(
  usuario_id BIGINT UNSIGNED NOT NULL,
  rol_id BIGINT UNSIGNED NOT NULL,
  FOREIGN KEY(usuario_id) REFERENCES users(id),
  FOREIGN KEY(rol_id) REFERENCES rol(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

LOCK TABLES `user_rol` WRITE;

UNLOCK TABLES;