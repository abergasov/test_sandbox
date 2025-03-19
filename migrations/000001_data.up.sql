create table observe_chats
(
    chat_id   BIGINT PRIMARY KEY,
    chat_name VARCHAR,
    chat_nick VARCHAR,
    active    BOOLEAN
);

create table chat_messages
(
    chat_id        BIGINT,
    message_id     BIGINT,
    timestamp      TIMESTAMP,
    timestamp_num  BIGINT,
    user_id        BIGINT,
    message        VARCHAR,
    reply_to       BIGINT,
    is_bot         BOOLEAN,
    is_deleted     BOOLEAN,
    is_edited      BOOLEAN
);

ALTER TABLE chat_messages ADD CONSTRAINT chat_messages_pk PRIMARY KEY (chat_id, message_id);
CREATE INDEX chat_messages_timestamp_num_index ON chat_messages (timestamp_num);
