CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE TABLE tracks (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    tags TEXT[],
    description TEXT
);

CREATE TABLE stages (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    track_id INTEGER NOT NULL REFERENCES tracks(id),
    description TEXT NOT NULL,
    level INTEGER NOT NULL
);

CREATE TABLE assessments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    reviewer_id INTEGER NOT NULL REFERENCES users(id),
    state TEXT DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE TABLE scores (
    assessment_id INTEGER NOT NULL REFERENCES assessments(id),
    track_id INTEGER NOT NULL REFERENCES tracks(id),
    stage_id INTEGER NOT NULL REFERENCES stages(id),
    score INTEGER NOT NULL,
    PRIMARY KEY (assessment_id, track_id, stage_id)
);

CREATE TABLE feedback (
    track_id INTEGER NOT NULL REFERENCES tracks(id),
    assessment_id INTEGER NOT NULL REFERENCES assessments(id),
    feedback TEXT,
    examples TEXT,
    recommendations TEXT,
    PRIMARY KEY (track_id, assessment_id)
);