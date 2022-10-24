CREATE TABLE SocialMedias (
    id SERIAL PRIMARY KEY NOT NULL, 
    name VARCHAR(255) NOT NULL, 
    social_media_url text NOT NULL, 
    user_id INT REFERENCES Users (id) ON DELETE CASCADE NOT NULL ,
    created_at DATE DEFAULt NOW()::date NOT NULL, 
    updated_at DATE DEFAULt NOW()::date NOT NULL 
);