CREATE TYPE task_status_enum AS ENUM ('PENDING', 'DONE', 'FAILED');

CREATE TABLE IF NOT EXISTS tasks(
    id serial PRIMARY KEY,
    created_by  INT NOT NULL REFERENCES users (id),
    title VARCHAR(300) NOT NULL,
    description TEXT,
    deleted boolean DEFAULT FALSE,
    status task_status_enum DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);