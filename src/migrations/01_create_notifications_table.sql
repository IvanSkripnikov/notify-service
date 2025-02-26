CREATE TABLE IF NOT EXISTS notifications (
    id INT auto_increment PRIMARY KEY,
    user_id INT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created BIGINT UNSIGNED
);