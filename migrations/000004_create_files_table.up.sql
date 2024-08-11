CREATE TABLE IF NOT EXISTS files
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT                   NOT NULL,
    file_url    VARCHAR(255)             NOT NULL,
    uploaded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_files_user_id ON files (user_id);
CREATE INDEX IF NOT EXISTS idx_files_uploaded_at ON files (uploaded_at);

ALTER TABLE files
ADD CONSTRAINT fk_files_user_id
FOREIGN KEY (user_id) REFERENCES users (id);