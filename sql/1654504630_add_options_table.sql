CREATE TABLE IF NOT EXISTS options (
  id uuid PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  poll_id uuid REFERENCES polls(id),
  votes INTEGER NOT NULL
);
