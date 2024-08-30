-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd