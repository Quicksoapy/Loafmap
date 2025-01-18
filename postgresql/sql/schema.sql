CREATE TABLE IF NOT EXISTS markers (
    id bytea PRIMARY KEY,
    userid string,
    description VARCHAR (520) NOT NULL,
    latitude VARCHAR (50) NOT NULL,
    longitude VARCHAR (50) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    image string
);