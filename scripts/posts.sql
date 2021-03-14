CREATE TABLE posts (
    post_id uuid DEFAULT uuid_generate_v4 (),
    title VARCHAR NOT NULL,
    body VARCHAR NOT NULL,
    id uuid references users(id),
    PRIMARY KEY (post_id)
);