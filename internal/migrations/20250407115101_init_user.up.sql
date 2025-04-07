CREATE TYPE role AS ENUM('admin' , 'user');
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    id_token TEXT,
    user_role  role
);