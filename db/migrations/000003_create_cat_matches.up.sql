CREATE TABLE cat_matches (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    match_cat_id SERIAL NOT NULL,
    user_cat_id SERIAL NOT NULL,
    message VARCHAR(120) NOT NULL,
    is_approved BOOLEAN NOT NULL default false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_cat_matches_all_columns ON cat_matches (user_cat_id, match_cat_id);