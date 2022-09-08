--Filename: 000001_kriol.up.sql


CREATE TABLE IF NOT EXISTS entries(

    id bigserial PRIMARY KEY,
    english_word VARCHAR ( 100 ) UNIQUE NOT NULL,
    kriol_word VARCHAR ( 100 ) UNIQUE NOT NULL
);