DROP DATABASE IF EXISTS `mercado-fresh-panic`;

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
  minimum_capacity BIGINT NOT NULL,
  minimum_temperature DECIMAL(19, 2) NOT NULL,
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
  due_date VARCHAR(255) NOT NULL,
  initial_quantity BIGINT UNSIGNED NOT NULL,
  manufacturing_date VARCHAR(255) NOT NULL,
  manufacturing_hour VARCHAR(255) NOT NULL,
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

CREATE TABLE `order_details`(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    clean_liness_status VARCHAR(255) NOT NULL,
    quantity BIGINT NOT NULL,
    temperature DECIMAL(19, 2) NOT NULL,
    product_record_id BIGINT UNSIGNED NOT NULL,
    purchase_order_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (product_record_id) REFERENCES product_records(id),
    FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `order_details` WRITE;

UNLOCK TABLES;

INSERT INTO countries(country_name) 
VALUES  ("Brasil"),
        ("Argentina");

INSERT INTO provinces(province_name, id_country_fk) 
VALUES  ("São Paulo", 1),
        ("Rio de Janeiro", 1),
        ("Minas Gerais", 1),
        ("Buenas Aires", 2),
        ("Formosa", 2);

INSERT INTO localities(id, locality_name, province_id) 
VALUES  ("11065001", "Santos", 1),
        ("10235001", "Campinas", 1),
        ("16372001", "Belo Horizonte", 3),
        ("11223001", "Buenos Aires", 4);

INSERT INTO sellers(cid, company_name, address, telephone, locality_id) 
VALUES  (1, "Nike", "Rua Pedro Américo, 212", "13990984533", "11065001"),
        (2, "Adidas", "Rua Washington Luis, 2212", "12990123453", "11065001"),
        (3, "Puma", "Rua Goiás, 2645", "15923484533", "11223001"),
        (4, "Multilaser", "Rua Maranhão, 5467", "13995325533", "11065001"),
        (5, "Logitech", "Rua Cafú, 1231", "13990235533", "16372001"),
        (6, "Hering", "Rua Pelé, 5467", "13990986533", "10235001");

INSERT INTO carriers(cid, company_name, address, telephone, locality_id) 
VALUES  (1, "Loggi", "Rua São Paulo, 212", "1332234533", "11065001"),
        (2, "DHL", "Rua Espírito Santo, 2212", "1232113453", "11065001");

INSERT INTO warehouses(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) 
VALUES  ("Rua São Paulo, 212", "DHH", "1332234533", 25000, 10.5, "11065001"),
        ("Rua Espírito Santo, 2212", "CJJ", "1232113453", 50000, 8, "11065001");

INSERT INTO products_types(description) 
VALUES  ("Roupa"),
        ("Eletrônico"),
        ("Cozinha");

INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type, seller_id)
VALUES  ("Calça jeans", 1, 1, 123, 60, 50, "CJ00", 12, 25, 1, 1),
        ("Mouse", 3, 6, 13, 20, 10, "M00", 12, 25, 2, 5),
        ("Tênis", 4, 21, 1223, 60, 50, "T00", 12, 25, 1, 3);

INSERT INTO sections(section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type, warehouse_id) 
VALUES	(1, 1000, 18, 10000, 500, 9, 2, 1),
		    (2, 2000, 12, 20000, 1500, 5, 1, 1),
		    (3, 3000, 28, 30000, 1500, 2, 1, 2);

INSERT INTO product_batches(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id)
VALUES	(1, 100, 12, "2022-08-23 13:15:00", 50, "2021-01-20", "12:11:00", 6, 2, 1),
		    (2, 200, 24, "2022-07-23 14:15:00", 100, "2021-01-20", "13:11:00", 26, 1, 2),
		    (3, 300, 36, "2022-09-23 15:15:00", 150, "2021-01-20", "15:11:00", 16, 2, 2);

INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id)
VALUES	("2022-07-04 09:16:12", 15.5, 10.5, 1),
		    ("2022-06-14 06:36:14", 55.90, 40.5, 2),
		    ("2022-05-24 07:15:32", 85.25, 50.5, 3);

INSERT INTO employees(id_card_number, first_name, last_name, warehouse_id)
VALUES	("1111222233334444", "José", "Neto", 1),
        ("1456542642455555", "Fernando", "Diniz", 1),
        ("2543542532354543", "Paulo", "Souza", 2);

INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id)
VALUES	("2022-03-21 12:11:21", "1234", 1, 1, 1),
		    ("2022-04-21 13:11:21", "2134", 1, 2, 2),
        ("2022-05-21 14:11:21", "3543", 2, 2, 2),
        ("2022-06-21 15:11:21", "3561", 2, 1, 1);
        
INSERT INTO order_status(description)
VALUES	("Aprovado"),
        ("Em trânsito"),
        ("Reprovado");

INSERT INTO buyers(id_card_number, first_name, last_name)
VALUES	("1111666633339999", "Carlos", "Neto"),
        ("1487565842455585", "Luiz", "Mendes"),
        ("2543544747425879", "Isabelle", "Silva");
        
INSERT INTO purchase_orders(order_number, order_date, tracking_code, buyer_id, order_status_id, product_record_id)
VALUES	("1234", "2021-02-27 18:11:32", "ABCD", 1, 1, 1),
		    ("5678", "2022-06-22 08:51:51", "EFGH", 1, 2, 3),
        ("9814", "2022-04-17 09:14:14", "GHIJ", 2, 1, 2);

INSERT INTO order_details(clean_liness_status, quantity, temperature, product_record_id, purchase_order_id)
VALUES  ("Aprovado", 1000, 12.5, 1, 1),
        ("Aprovado", 2000, 11.5, 2, 2),
        ("Rejeitado", 3000, 10.5, 1, 2);