-- +migrate Up
CREATE TYPE task_status AS ENUM ('todo', 'in_progress', 'done');

CREATE TABLE tasks (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    status      task_status  NOT NULL DEFAULT 'todo',
    assignee    VARCHAR(255),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS tasks;
DROP TYPE  IF EXISTS task_status;

