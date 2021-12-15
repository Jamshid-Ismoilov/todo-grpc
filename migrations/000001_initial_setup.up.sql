CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL Primary Key,
    assignee varchar(64),
    title varchar(64),
    summary varchar(64),
    deadline timestamp default current_timestamp,
    status varchar(64),
);