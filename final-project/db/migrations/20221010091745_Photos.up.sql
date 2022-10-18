CREATE TABLE Photos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL, 
    caption VARCHAR(255) NOT NULL, 
    photo_url VARCHAR(255) NOT NULL, 
    user_id INT REFERENCES Users (id) NOT NULL, 
    created_at TIMESTAMP DEFAULt NOW() NOT NULL, 
    updated_at TIMESTAMP DEFAULt NOW() NOT NULL 
);