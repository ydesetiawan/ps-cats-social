CREATE TABLE users (
    id bigserial PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    password VARCHAR(100)
);


CREATE INDEX
    user_index_1 ON users (name, email);
