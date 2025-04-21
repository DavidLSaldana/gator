-- +goose Up
CREATE TABLE posts(
	id INT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
	feed_id INT NOT NULL REFERENCES feeds(id)
		ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
