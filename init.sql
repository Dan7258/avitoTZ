
CREATE TABLE IF NOT EXISTS users (
    id        VARCHAR(255) NOT NULL PRIMARY KEY,
    username  VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    team_name VARCHAR(255) NOT NULL

);

CREATE TABLE IF NOT EXISTS teams (
     team_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS pull_request (
    pull_request_id VARCHAR(255) NOT NULL PRIMARY KEY,
    pull_request_name VARCHAR(255) NOT NULL,
    author_id VARCHAR(255) NOT NULL,
    status VARCHAR(6) NOT NULL CHECK ( status IN ('OPEN', 'MERGED') ),
    assigned_reviews TEXT[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    merged_at TIMESTAMPTZ NULL
);


