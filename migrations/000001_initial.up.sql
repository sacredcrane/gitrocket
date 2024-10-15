CREATE DOMAIN defaulted_timestamp AS TIMESTAMP DEFAULT NOW();

CREATE TYPE status_e AS ENUM ('READY', 'WORKING', 'PENDING', 'ERRORED');

CREATE TYPE event_type_e AS ENUM ('ON_MERGE', 'ON_TAG', 'ON_COMMIT');

CREATE TYPE rule_type_e AS ENUM ('BY_TIME', 'BY_EVENT');

CREATE TABLE IF NOT EXISTS integration_rule
(
    rule_id uuid DEFAULT gen_random_uuid(),
    rule_type rule_type_e NOT NULL DEFAULT 'BY_TIME',
    time_rule INTERVAL,
    event_type event_type_e,
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    PRIMARY KEY(rule_id)
);

CREATE TABLE IF NOT EXISTS project_settings
(
    settings_id uuid DEFAULT gen_random_uuid(),
    settings_rpath VARCHAR(1024) NOT NULL,
    settings_file bytea NOT NULL,
    dockerfile_rpath VARCHAR(1024),
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    PRIMARY KEY(settings_id)
);

CREATE TABLE IF NOT EXISTS docker_host
(
    dhost_id uuid DEFAULT gen_random_uuid(),
    docker_user VARCHAR(128) NOT NULL,
    host_ip inet NOT NULL,
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    last_access_at defaulted_timestamp,
    PRIMARY KEY(dhost_id) 
);

CREATE TABLE IF NOT EXISTS git_repo
(
    repo_id uuid DEFAULT gen_random_uuid(),
    repo_url VARCHAR(2048) NOT NULL,
    tracked_branch VARCHAR(128) NOT NULL,
    access_token VARCHAR(1024),
    last_access_at TIMESTAMP,
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    PRIMARY KEY(repo_id)
);

CREATE TABLE IF NOT EXISTS gr_user
(
    gr_user_id uuid DEFAULT gen_random_uuid(),
    user_login VARCHAR(128) NOT NULL,
    user_password VARCHAR(128) NOT NULL,
    last_access_at defaulted_timestamp,
    access_token VARCHAR(1024),
    last_used_at TIMESTAMP,
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    PRIMARY KEY(gr_user_id)
);

CREATE TABLE IF NOT EXISTS project
(
    project_id uuid DEFAULT gen_random_uuid(),
    owner_id uuid NOT NULL,
    repo_id uuid NOT NULL,
    settings_id uuid NOT NULL,
    dhost_id uuid NOT NULL,
    rule_id uuid NOT NULL,
    project_status status_e DEFAULT 'READY',
    created_at defaulted_timestamp,
    updated_at defaulted_timestamp,
    project_desc TEXT,
    PRIMARY KEY(project_id),
    FOREIGN KEY(owner_id) REFERENCES gr_user(gr_user_id),
    FOREIGN KEY(repo_id) REFERENCES git_repo(repo_id),
    FOREIGN KEY(settings_id) REFERENCES project_settings(settings_id),
    FOREIGN KEY(dhost_id) REFERENCES docker_host(dhost_id),
    FOREIGN KEY(rule_id) REFERENCES integration_rule(rule_id)
);