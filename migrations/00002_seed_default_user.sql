-- +goose Up
INSERT INTO schools (
  id,
  school_id,
  name,
  sub_domain,
  is_active,
  created_at,
  updated_at,
  created_by,
  updated_by
) VALUES (
  'default-school',
  'default-school',
  'Default School',
  'default-school',
  true,
  now(),
  now(),
  'system',
  'system'
);

INSERT INTO users (
  id,
  school_id,
  email,
  full_name,
  password_hash,
  role_id,
  status,
  created_at,
  updated_at,
  created_by,
  updated_by
) VALUES (
  'user-1',
  'default-school',
  'test@school.org',
  'Test Admin',
  '$2b$12$nLwMyOMzrftAUnbzJqfOiOim7b1zmcnpPY0mDrnua.rKSbTsDhffG',
  'admin',
  'active',
  now(),
  now(),
  'system',
  'system'
);

-- +goose Down
DELETE FROM users WHERE id = 'user-1';
DELETE FROM schools WHERE id = 'default-school';
