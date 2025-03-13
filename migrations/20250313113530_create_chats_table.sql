-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats
(
    id bigserial primary key,
    created_at timestamp(0) default CURRENT_TIMESTAMP,
    updated_at timestamp(0)
);
CREATE TABLE chats_users
(
    id bigserial primary key,
    chat_id int references chats(id) on delete cascade,
    username varchar(64) not null,
    created_at timestamp(0) default CURRENT_TIMESTAMP,
    updated_at timestamp(0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats_users;
DROP TABLE chats;
-- +goose StatementEnd
