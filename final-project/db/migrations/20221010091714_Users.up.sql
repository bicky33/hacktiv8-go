CREATE TABLE Users (
    id SERIAL PRIMARY KEY NOT NULL, 
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL, 
    profile_image_url text NOT NULL,
    age INT NOT NULL, 
    created_at DATE DEFAULt NOW()::date NOT NULL, 
    updated_at DATE DEFAULt NOW()::date NOT NULL 
);