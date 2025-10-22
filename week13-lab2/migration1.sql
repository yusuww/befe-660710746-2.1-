-- 1. Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    email_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- Index สำหรับ login
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_active ON users(is_active);

-- Admin user (password: admin123)
INSERT INTO users (username, email, password_hash, email_verified)
VALUES (
    'admin',
    'admin@bookstore.com',
    '$2a$12$BsZ.rdXlcqpNKuispo6aYegCa8NM1Bbg.Lb.RhBaOWJ.SG5yunhte',
    true
);

-- Editor user (password: editor123)
INSERT INTO users (username, email, password_hash, email_verified)
VALUES (
    'poohkan',
    'editor@bookstore.com',
    '$2a$12$uc5C2LrsKKwZ/K4wim.RN.VjsKsClMwL9wHHVwEE3fFgYEvCrfpz',
    true
);

-- Regular user (password: user123)
INSERT INTO users (username, email, password_hash, email_verified)
VALUES (
    'nuttachot',
    'user@bookstore.com',
    '$2a$12$POzHIJhq6O3VaawJQDsL1uOoADykpywnAO1ecRlwccC7glIf0kpCW',
    true
);