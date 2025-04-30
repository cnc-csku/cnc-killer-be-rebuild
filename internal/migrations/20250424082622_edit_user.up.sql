ALTER TABLE users
DROP COLUMN id_token,
ADD COLUMN refresh_token TEXT,
ADD COLUMN email TEXT;

-- change primary key
ALTER TABLE users DROP CONSTRAINT users_pkey;
ALTER TABLE users DROP COLUMN user_id;
ALTER TABLE users ADD PRIMARY KEY (email);