CREATE TABLE IF NOT EXISTS tasks(
    id varchar(256),
    assignee varchar(64),
    title varchar(64),
    summary varchar(64),
    deadline timestamp default current_timestamp,
    status bool,
    created_at timestamp default current_timestamp,
    updated_at timestamp default null,
    deleted_at timestamp default null
);