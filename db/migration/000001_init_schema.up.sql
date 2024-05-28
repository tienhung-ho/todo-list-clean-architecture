CREATE TYPE status_enum AS ENUM ('Doing', 'Done', 'Deleted');

CREATE TABLE todo_items (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  image JSON,
  description TEXT,
  status status_enum  DEFAULT 'Doing',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
