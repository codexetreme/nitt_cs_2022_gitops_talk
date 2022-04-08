CREATE TABLE IF NOT EXISTS roster (
    id VARCHAR (9) PRIMARY KEY,
    name VARCHAR (50)
);

INSERT INTO roster(id,name) VALUES ('0000','admin') RETURNING id;