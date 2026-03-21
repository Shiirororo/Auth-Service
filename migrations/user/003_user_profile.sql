CREATE TABLE user_profile (
    user_id BINARY(16) NOT NULL,

    profile_name VARCHAR(50) NOT NULL,
    mobile VARCHAR(20),
    gender TINYINT,
    birthday DATE,

    PRIMARY KEY (user_id),

    CONSTRAINT fk_profile_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_profile_name
ON user_profile(profile_name);