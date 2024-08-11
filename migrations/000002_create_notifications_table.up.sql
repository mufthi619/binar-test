CREATE TABLE IF NOT EXISTS notifications
(
    id      BIGSERIAL PRIMARY KEY,
    user_id BIGINT                   NOT NULL,
    message VARCHAR(255)             NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications (user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_sent_at ON notifications (sent_at);