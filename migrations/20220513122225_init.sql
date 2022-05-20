-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id         bigint          PRIMARY KEY NOT NULL,
    username        varchar(255)    NOT NULL DEFAULT '',
    correct_answers int             NOT NULL DEFAULT 0,
    total_answers   int             NOT NULL DEFAULT 0,
    is_passing      boolean         NOT NULL DEFAULT false,
    registered_at   timestamp       NOT NULL DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
