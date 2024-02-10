CREATE TABLE IF NOT EXISTS profiles (
   id bigserial PRIMARY KEY,
   username VARCHAR (50) NOT NULL,
   firstname VARCHAR (50) NOT NULL,
   lastname VARCHAR (50) NOT NULL,
   email VARCHAR (300) NOT NULL,
   phone VARCHAR (12) NOT NULL
);