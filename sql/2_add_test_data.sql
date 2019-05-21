-- set up tracks
INSERT INTO tracks (name, tags) VALUES (
    'Coding',
    '{"engineer"}'
);

INSERT INTO tracks (name, tags) VALUES (
    'Kaizen',
    '{"all"}'
);

-- set up stages for Coding track
INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Coding'),
    'Contributes to project development. Understands coding best practices as set forth by the team',
    1
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Coding'),
    'Develops features from design through implementation. Writes code that is consistent and follows best practices. Capable of debugging most common errors',
    2
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Coding'),
    'Develops products from design through implementation. Writes code that is correct, efficient, and easy to understand. Excellent at debugging and fixing errors, using tests where appropriate',
    3
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Coding'),
    'Develops code that is modular and reusable across projects. Writes code that is performant. Knows how and when to profile, identify, and optimize performance bottlenecks',
    4
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Coding'),
    'Achieves simplicity over complexity in code. Develops platforms and tools that enable the rapid development of new projects',
    5
);

-- set up stages for Kaizen track
INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Kaizen'),
    'Curious and passionate. Places a strong focus on learning and self-improvement',
    1
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Kaizen'),
    'Keeps abreast of relevant trends. Connects learnings to improvements in their day-to-day role. Continues to learn and improve through iteration',
    2
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Kaizen'),
    'Always exploring and testing new ideas. Sets and follows through on self-improvement goals, prioritizing areas that create maximum ROI',
    3
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Kaizen'),
    'Sets and follows through on team goals. Rigorously strategizes, prioritizes, and pushes forward areas of improvement that maximize team ROI. Proactively looks for and suggests areas for improvement, putting wheels in motion independently where possible',
    4
);

INSERT INTO stages (track_id, description, level) VALUES (
    (SELECT id FROM tracks where name='Kaizen'),
    'Designs improvement plans and sets targets and priorities for teams and individuals. Productively ""plugged in"" to relevant areas in the industry, helping draw from best-in-class practices to improve the discipline company-wide',
    5
);

-- create users
INSERT INTO users (name, email, hashed_password, tags) VALUES (
    'Bruno Bachmann',
    'bruno.bachmann@dapperlabs.com',
    '$2a$04$lpV3TloLJSLyD9rpbwDcueQeRvD4JZqsZVHUbeE8UbOqasUQOZpp.',
    '{"engineer","all"}'
);

INSERT INTO users (name, email, hashed_password, tags) VALUES (
    'Brant Hardy',
    'brant@dapperlabs.com',
    '$2a$04$lpV3TloLJSLyD9rpbwDcueQeRvD4JZqsZVHUbeE8UbOqasUQOZpp.',
    '{"engineer","all"}'
);

-- create an assessment for the user
INSERT INTO assessments (user_id, reviewer_id, state) VALUES (
    (SELECT id FROM users WHERE name='Bruno Bachmann'),
    (SELECT id FROM users WHERE name='Brant Hardy'),
    'complete'
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    1,
    (SELECT id FROM tracks where name='Coding'),
    1,
    4
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    1,
    (SELECT id FROM tracks where name='Coding'),
    2,
    2
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    1,
    (SELECT id FROM tracks where name='Kaizen'),
    1,
    2
);

INSERT INTO feedback (track_id, assessment_id, feedback, examples, recommendations) VALUES (
    (SELECT id FROM tracks WHERE name='Coding'),
    (SELECT id FROM assessments WHERE state='complete'),
    'Needs a lot of work',
    'You wrote a terrible bug today',
    'Learn how to code, buddy'
);

-- create a second assessment for the same user
INSERT INTO assessments (user_id, reviewer_id, state) VALUES (
    (SELECT id FROM users WHERE name='Bruno Bachmann'),
    (SELECT id FROM users WHERE name='Brant Hardy'),
    'pending'
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    2,
    (SELECT id FROM tracks where name='Coding'),
    1,
    5
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    2,
    (SELECT id FROM tracks where name='Coding'),
    2,
    3
);

INSERT INTO scores (assessment_id, track_id, stage_id, score) VALUES (
    2,
    (SELECT id FROM tracks where name='Kaizen'),
    1,
    4
);

INSERT INTO feedback (track_id, assessment_id, feedback, examples, recommendations) VALUES (
    (SELECT id FROM tracks WHERE name='Coding'),
    (SELECT id FROM assessments WHERE state='pending'),
    'Improved a lot',
    'You wrote a good feature',
    'Keep practicing'
);