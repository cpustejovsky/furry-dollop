CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 (),
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    expertise VARCHAR NOT NULL,
    hashed_password VARCHAR NOT NULL,
    PRIMARY KEY (id)
);

