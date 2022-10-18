CREATE TABLE SocialMedias (
    id SERIAL PRIMARY KEY NOT NULL, 
    name VARCHAR(255) NOT NULL, 
    social_media_url text NOT NULL, 
    user_id INT REFERENCES Users (id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);