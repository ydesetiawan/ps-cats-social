CREATE TABLE users (
                       id bigserial PRIMARY KEY,
                       name VARCHAR(50),
                       email VARCHAR(50),
                       password VARCHAR(50)
);


CREATE INDEX user_index_1 ON users (name, email);
