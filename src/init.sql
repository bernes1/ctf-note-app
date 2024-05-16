CREATE TABLE Platform (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE Artist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE DJset (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    platform_id INTEGER NOT NULL,
    artist_id INTEGER NOT NULL,
    FOREIGN KEY (platform_id) REFERENCES Platform(id),
    FOREIGN KEY (artist_id) REFERENCES Artist(id)
);
