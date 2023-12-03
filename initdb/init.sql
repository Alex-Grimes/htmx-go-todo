CREATE TABLE
    todos (
        id SERIAL NOT NULL PRIMARY KEY,
        todo VARCHAR(255) NOT NULL,
        done BOOLEAN
    );

INSERT INTO todos (todo, done) VALUES ('Hello World!', false);