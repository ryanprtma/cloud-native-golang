ALTER TABLE users
ADD deleted_at timestamp DEFAULT (now()) AFTER created_at;