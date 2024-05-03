CREATE TYPE race_enum AS ENUM ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman');
CREATE TYPE sex_enum AS ENUM ('male', 'female');

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    race race_enum NOT NULL,
    sex sex_enum NOT NULL,
    age_in_month INT NOT NULL,
    image_urls TEXT[],
    description VARCHAR(200) NOT NULL,
    has_matched BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX idx_cats_all_columns ON cats (name);