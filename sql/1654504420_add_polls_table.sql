CREATE TABLE IF NOT EXISTS polls (
  id uuid PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL
);