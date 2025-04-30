ALTER TABLE users
ADD COLUMN id_token TEXT,
DROP COLUMN refresh_token,
ADD COLUMN user_id UUID;

-- for droping current pk
ALTER TABLE users DROP CONSTRAINT users_pkey;
ALTER TABLE users DROP COLUMN email;
ALTER TABLE users ADD PRIMARY KEY (user_id);