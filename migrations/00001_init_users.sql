-- +goose Up
CREATE TABLE IF NOT EXISTS schools (
  id text PRIMARY KEY,
  school_id text NOT NULL,
  name text NOT NULL,
  sub_domain text NOT NULL,
  is_active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL,
  created_by text NOT NULL,
  updated_by text NOT NULL,
  deleted_at timestamptz NULL,
  deleted_by text NULL
);

CREATE TABLE IF NOT EXISTS users (
  id text PRIMARY KEY,
  school_id text NOT NULL,
  email text NOT NULL UNIQUE,
  full_name text NOT NULL,
  password_hash text NOT NULL,
  role_id text NOT NULL,
  status text NOT NULL,
  last_login_at timestamptz NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL,
  created_by text NOT NULL,
  updated_by text NOT NULL,
  deleted_at timestamptz NULL,
  deleted_by text NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS schools;
