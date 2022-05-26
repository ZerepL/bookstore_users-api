CREATE DATABASE users_db;

CREATE TABLE users_db.users (
    id           INT NOT NULL AUTO_INCREMENT,
    first_name   VARCHAR(45) NOT NULL,
    last_name    VARCHAR(45) NOT NULL,
    email 		 VARCHAR(45) NOT NULL,
    date_created DATETIME NOT NULL,
    status 		 VARCHAR(45) NOT NULL,
    password     VARCHAR(32) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX email_UNIQUE (email ASC)
);