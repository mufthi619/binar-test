CREATE TABLE messages
(
    id               BIGSERIAL PRIMARY KEY,
    conversation_id  BIGINT NOT NULL,
    sender_id        BIGINT NOT NULL,
    content          TEXT NOT NULL,
    sent_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP WITH TIME ZONE
);

CREATE TABLE conversations
(
    id                   BIGSERIAL PRIMARY KEY,
    participant_user_ids BIGINT[] NOT NULL DEFAULT '{}'::BIGINT[],
    created_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at           TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_messages_conversation_id ON messages (conversation_id);
CREATE INDEX idx_messages_sender_id ON messages (sender_id);
CREATE INDEX idx_messages_sent_at ON messages (sent_at);
CREATE INDEX idx_messages_deleted_at ON messages (deleted_at);
CREATE INDEX idx_conversations_deleted_at ON conversations (deleted_at);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_messages_updated_at
    BEFORE UPDATE ON messages
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_conversations_updated_at
    BEFORE UPDATE ON conversations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();