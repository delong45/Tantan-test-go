
CREATE TABLE users (id serial primary key, name text, "type" text);
CREATE TABLE relationships (id text, other_id text, state text);
