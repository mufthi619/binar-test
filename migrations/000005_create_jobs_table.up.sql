CREATE TABLE IF NOT EXISTS jobs
(
    id           BIGSERIAL PRIMARY KEY,
    status       VARCHAR(20)  NOT NULL,
    queued_at    TIMESTAMP    NOT NULL,
    completed_at TIMESTAMP,
    message      VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_jobs_status ON jobs (status);
CREATE INDEX IF NOT EXISTS idx_jobs_queued_at ON jobs (queued_at);
CREATE INDEX IF NOT EXISTS idx_jobs_completed_at ON jobs (completed_at);