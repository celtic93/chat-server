-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    id bigserial primary key,
    chat_id int references chats(id) on delete cascade,
    user_id int not null,
    text text not null,
    created_at timestamp(0) not null default CURRENT_TIMESTAMP,
    updated_at timestamp(0) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
