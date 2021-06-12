-- +goose Up
-- SQL in this section is executed when the migration is applied.
create schema if not exists account;

CREATE TABLE IF NOT EXISTS account.users(
    id serial NOT NULL,
    uuid uuid DEFAULT uuid_generate_v4 () UNIQUE,
    created_at            timestamptz NOT NULL DEFAULT now(),
    updated_at            timestamptz NOT NULL DEFAULT now(),
    first_name varchar not null,
    middle_name varchar not null default '',
    last_name varchar not null,
    login varchar not null UNIQUE,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
alter table account.users add column if not exists confirmed boolean not null default false;
alter table account.users add column if not exists avatar_url varchar not null default '';

DROP TRIGGER IF EXISTS set_timestamp on account.users;
create trigger set_timestamp before update
    on
    account.users for each row execute procedure trigger_set_timestamp();


CREATE TABLE IF NOT EXISTS account.login_history(
    id serial NOT NULL,
    created_at            timestamptz NOT NULL DEFAULT now(),
    user_id int not null REFERENCES account.users(id),
    ip varchar not null default '',
    success boolean not null default false,
    
    CONSTRAINT login_history_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS account.passwords(
    id serial NOT NULL,
    created_at            timestamptz NOT NULL DEFAULT now(),
    user_id int not null REFERENCES account.users(id),
    pass varchar not null,
    
    CONSTRAINT passwords_pkey PRIMARY KEY (id)
);

create unique index if not exists login_users_uidx on account.users (login)
WHERE login != '';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.