CREATE TABLE IF NOT EXISTS accounts (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (520) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS markers (
    id serial PRIMARY KEY,
    userid INT,
    description VARCHAR (520) NOT NULL,
    latitude VARCHAR (50) NOT NULL,
    longitude VARCHAR (50) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    imageid INT
);