CREATE TYPE USER_ROLE AS ENUM ('super_admin', 'owner', 't3_admin', 'admin', 'user', 'viewer', 'support');

CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR(50) NOT NULL UNIQUE,
    password    text        NOT NULL,
    first_name  VARCHAR(50) NOT NULL,
    last_name   VARCHAR(50) NOT NULL,
    time_zone   VARCHAR(50) DEFAULT NULL,
    mobile      VARCHAR(20) DEFAULT NULL,
    role        USER_ROLE   DEFAULT 'user',
    is_active   BOOLEAN     DEFAULT TRUE,
    created_at  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_id_idx ON users (id);