CREATE TYPE user_type_enum AS ENUM ('ADMIN','USER','REVIEWER');

CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   fullname VARCHAR (225) NOT NULL,
   password VARCHAR (225) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   user_type user_type_enum DEFAULT 'USER',
   deleted boolean DEFAULT FALSE,
   created_by INT REFERENCES users (id),
   created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
   updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);
