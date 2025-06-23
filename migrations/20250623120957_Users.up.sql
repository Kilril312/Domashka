CREATE TABLE Users (
                                    id SERIAL PRIMARY KEY,
                                    Email VARCHAR(255) NOT NULL,
                                    Password VARCHAR(255) NOT NULL,
                                    CreatedAt TIMESTAMP NOT NULL DEFAULT NOW(),
                                    UpdatedAt TIMESTAMP NOT NULL DEFAULT NOW(),
                                    DeletedAt TIMESTAMP DEFAULT NULL
);