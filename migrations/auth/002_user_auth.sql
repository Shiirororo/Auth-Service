CREATE TABLE user_auth (
    user_id BINARY(16) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password_hash VARCHAR(72) NOT NULL,

    last_login TIMESTAMP NULL,
    locked_until TIMESTAMP NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id),

    CONSTRAINT fk_user_auth_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE UNIQUE INDEX idx_user_auth_email
ON user_auth(email);