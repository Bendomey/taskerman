-- CREATE TYPE review_status_enum AS ENUM ('PENDING','ACCEPT','REJECT');

CREATE TABLE IF NOT EXISTS reviews(
    id serial PRIMARY KEY,
    task_id serial NOT NULL REFERENCES tasks (id),
    reviewer serial NOT NULL REFERENCES users (id),
    status review_status_enum DEFAULT 'PENDING',
    reason TEXT,
    deleted boolean DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);