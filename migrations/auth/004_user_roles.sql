CREATE TABLE roles (
    id SMALLINT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY idx_roles_name (name)
) ENGINE=InnoDB;