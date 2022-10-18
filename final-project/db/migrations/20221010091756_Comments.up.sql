CREATE TABLE Comments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users (id) NOT NULL, 
    photo_id INT REFERENCES Photos (id) NOT NULL,
    message TEXT NOT NULL, 
    created_at TIMESTAMP DEFAULt NOW() NOT NULL, 
    updated_at TIMESTAMP DEFAULt NOW() NOT NULL 
);