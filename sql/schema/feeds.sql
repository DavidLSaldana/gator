-- +goose Up
CREATE TABLE feeds(
	id INT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id INT NOT NULL REFERENCES users(id)
		ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;
