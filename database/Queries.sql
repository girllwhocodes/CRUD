CREATE TABLE users
(
    id         serial PRIMARY KEY,
    first_name VARCHAR(55),
    last_name  VARCHAR(55),
    login      VARCHAR(55),
    password   VARCHAR(55),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE educations
(
    id         serial PRIMARY KEY,
    title      VARCHAR(150) NOT NULL,
    country    VARCHAR(55) NOT NULL,
    date       DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE documents
(
    id         serial PRIMARY KEY,
    title   VARCHAR(77) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

select * from documents;

SELECT * FROM educations;

select  * from users;