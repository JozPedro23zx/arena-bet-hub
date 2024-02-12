CREATE TABLE tournaments
(
    id         TEXT NOT NULL,
    name       TEXT NOT NULL,
    event_date TEXT NOT NULL,
    street     TEXT NOT NULL,
    city       TEXT NOT NULL,
    state      TEXT NOT NULL,
    country    TEXT NOT NULL,
    finished   BOOLEAN NOT NULL
);