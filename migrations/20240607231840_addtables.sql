-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_info
(
    id         INTEGER,
    first_name TEXT   NOT NULL,
    last_name TEXT   NOT NULL,
    dateOfBirthday TIMESTAMP NOT NULL
);
CREATE TABLE users_credentials
(
    id         BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE plots
(
    id         BIGSERIAL PRIMARY KEY,
    user_id     integer not null,
    name        TEXT not null,
    content    jsonb not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE plots;
DROP TABLE users_info;
DROP TABLE users_credentials;
-- +goose StatementEnd
