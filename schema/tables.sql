CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE TABLE reviewers (
    reviewee_id NOT NULL REFERENCES users(id) ON DELETE CASCASE,
    reviewer_id NOT NULL REFERENCES users(id) ON DELETE CASCASE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE TABLE tracks (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE scores (
    user_id NOT NULL REFERENCES users(id) ON DELETE CASCASE,
    track_id NOT NULL REFERENCES tracks(id) ON DELETE CASCASE,
    stage INTEGER NOT NULL,
    score INTEGER NOT NULL
);

CREATE TABLE stages (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    description TEXT NOT NULL,
    level INTEGER NOT NULL
);