
-- tb_spu
CREATE TABLE `sd_product` (
    `productId` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `productName` VARCHAR(64) DEFAULT NULL COMMENT 'spu name',
    `productDesc` VARCHAR(256) DEFAULT NULL COMMENT 'spu desc',
    `productStatus` TINYINT(4) DEFAULT NULL COMMENT '0: out of stock, 1: in stock',
    `productAttribute` JSON DEFAULT NULL COMMENT 'json type attribute'
    `productShopID` BIGINT(20) DEFAULT NULL COMMENT 'shop id',
    `isDeleted` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '0: delete 1:null',
    `sort` INT(10) DEFAULT '0' COMMENT 'priority sort',
    `createTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'created timestamp', 
    `lastUpdate` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'last update timestamp',

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='spu';


-- tb_sku
CREATE TABLE `sku` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `skuNo` VARCHAR(32) DEFAULT '' COMMENT 'sku no',
    `skuName` VARCHAR(50) DEFAULT NULL COMMENT 'sku_name',
    `skuDesc` VARCHAR(256) DEFAULT NULL COMMENT 'sku desc',
    `skuType` TINYINT(4) DEFAULT NULL COMMENT 'sku_type',
    `status` TINYINT(4) NOT NULL COMMENT 'status',
    `sort` INT(10) DEFAULT '0' COMMENT 'priority sort',
    `skuStock` INT(11) NOT NULL DEFAULT '0' COMMENT 'sku stock',
    `skuPrice` DECIMAL(8,2) NOT NULL COMMENT 'sku price',
    `createTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'CREATED TIME',
    `lastUpdate` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'LAST UPDATE TIME',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_sku_no` (`skuNo`) USING BTREE
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = 'sku';

CREATE TABLE `sku_attr` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `skuNo` VARCHAR(32) DEFAULT '' COMMENT 'sku no',
    -- `skuStock` INT(11) NOT NULL DEFAULT '0' COMMENT 'sku stock',
    `skuAttribute` JSON DEFAULT NULL COMMENT 'sku attribute',
    `createTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'CREATED TIME',
    `lastUpdate` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'LAST UPDATE TIME',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_sku_no` (`skuNo`) USING BTREE
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = 'sku_attr';


CREATE TABLE `spu_to_sku` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `skuNo` VARCHAR(32) NOT NULL,
    `spuNo` VARCHAR(32) NOT NULL,

    `isDeleted` TINYINT UNSIGNED DEFAULT 0,
    `createTime` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `lastUpdate` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_spu_sku` (`spuNo`, `skuNo`),
    KEY `idx_sku` (`skuNo`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




-- CREATE TABLE products (
--     id BIGINT PRIMARY KEY AUTO_INCREMENT,
--     name VARCHAR(255) NOT NULL,
--     slug VARCHAR(255) UNIQUE,
--     description TEXT,
--     brand_id BIGINT,
--     category_id BIGINT,
    
--     price DECIMAL(12,2) NOT NULL,
--     original_price DECIMAL(12,2),

--     stock INT DEFAULT 0,
--     sold_count INT DEFAULT 0,

--     rating_avg FLOAT DEFAULT 0,
--     rating_count INT DEFAULT 0,

--     status ENUM('active', 'inactive', 'out_of_stock') DEFAULT 'active',

--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

--     INDEX idx_category (category_id),
--     INDEX idx_price (price),
--     INDEX idx_status (status),
--     INDEX idx_created (created_at)
-- );


