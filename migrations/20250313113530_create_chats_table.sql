-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats
(
    id bigserial primary key,
    created_at timestamp(0) not null default CURRENT_TIMESTAMP,
    updated_at timestamp(0) not null
);
CREATE TABLE chats_users
(
    id bigserial primary key,
    chat_id int references chats(id) on delete cascade,
    user_id int not null,
    created_at timestamp(0) not null default CURRENT_TIMESTAMP,
    updated_at timestamp(0) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats_users;
DROP TABLE chats;
-- +goose StatementEnd
