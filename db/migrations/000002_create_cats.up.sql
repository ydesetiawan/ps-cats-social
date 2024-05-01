CREATE TYPE race_enum AS ENUM ('Persian', 'MaineCoon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'BritishShorthair', 'Abyssinian', 'ScottishFold', 'Birman');
CREATE TYPE sex_enum AS ENUM ('male', 'female');

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    race race_enum NOT NULL,
    sex sex_enum NOT NULL,
    age_in_month INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    has_matched BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX idx_cats_all_columns ON cats (name);

CREATE TABLE cat_images (
    id SERIAL PRIMARY KEY,
    cat_id SERIAL NOT NULL,
    url TEXT NOT NULL,
    FOREIGN KEY (cat_id) REFERENCES cats(id)
);