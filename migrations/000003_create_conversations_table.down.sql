DROP TRIGGER IF EXISTS update_messages_updated_at ON messages;
DROP TRIGGER IF EXISTS update_conversations_updated_at ON conversations;

DROP TABLE IF EXISTS conversation_participants;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS conversations;

DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;