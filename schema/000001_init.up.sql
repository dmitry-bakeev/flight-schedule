CREATE TABLE IF NOT EXISTS flight
(
    id SERIAL PRIMARY KEY,
    number_flight INT NOT NULL,
    from_city VARCHAR(255) NOT NULL,
    time_from_city TIMESTAMP NOT NULL,
    to_city VARCHAR(255) NOT NULL,
    time_to_city TIMESTAMP NOT NULL
);
