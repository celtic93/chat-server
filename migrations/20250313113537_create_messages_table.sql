-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    id bigserial primary key,
    chat_id int references chats(id) on delete cascade,
    username varchar(64) not null,
    text text not null,
    created_at timestamp(0) default CURRENT_TIMESTAMP,
    updated_at timestamp(0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
