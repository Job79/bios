create table "user"
(
    id       serial                         not null,
    uid      uuid default gen_random_uuid() not null,
    name     text                           not null,
    password bytea                          not null,

    primary key (id),
    unique (uid),
    unique (name)
);

create table token
(
    token           bytea     not null,
    user_id         integer   not null,
    expiration_date timestamp not null,

    primary key (token),
    foreign key (user_id) references "user" on update cascade on delete cascade
);

create type movie_status as enum ('now', 'soon', 'ended');

create table movie
(
    id           serial                         not null,
    uid          uuid default gen_random_uuid() not null,
    name         text                           not null,
    description  text,
    total_time   integer                        not null,
    release_date timestamp                      not null,
    status       movie_status                   not null,

    primary key (id),
    unique (uid)
);

create index text_search on movie using gin(to_tsvector('english', name));

create table classification
(
    id          serial                         not null,
    uid         uuid default gen_random_uuid() not null,
    name        text                           not null,
    description text,

    primary key (id),
    unique (uid)
);

create table movie_classification
(
    movie_id          serial  not null,
    classification_id integer not null,

    primary key (movie_id, classification_id),
    foreign key (classification_id) references classification on update cascade on delete cascade,
    foreign key (movie_id) references movie on update cascade on delete cascade
);

create table genre
(
    id   serial                         not null,
    uid  uuid default gen_random_uuid() not null,
    name text                           not null,

    primary key (id),
    unique (uid)
);

create table movie_genre
(
    movie_id integer not null,
    genre_id integer not null,

    primary key (movie_id, genre_id),
    foreign key (movie_id) references movie on update cascade on delete cascade,
    foreign key (genre_id) references genre on update cascade on delete cascade
);

create table room
(
    id   serial                         not null,
    uid  uuid default gen_random_uuid() not null,
    code text                           not null,

    primary key (id),
    unique (uid)
);

create table showing
(
    id         serial                         not null,
    uid        uuid default gen_random_uuid() not null,
    movie_id   integer                        not null,
    room_id    integer                        not null,
    start_time timestamp                      not null,
    end_time   timestamp                      not null,

    primary key (id),
    unique (uid),
    foreign key (movie_id) references movie on update cascade on delete cascade,
    foreign key (room_id) references room on update cascade on delete cascade
);


create type file_type as enum ('image', 'video');

create table file
(
    id   serial                         not null,
    uid  uuid default gen_random_uuid() not null,
    path text                           not null,
    type file_type                      not null,

    primary key (id),
    unique (uid)
);

create table movie_file
(
    movie_id integer           not null,
    file_id  integer           not null,
    position integer default 0 not null,

    primary key (movie_id, file_id),
    foreign key (file_id) references file on update cascade on delete cascade,
    foreign key (movie_id) references movie on update cascade on delete cascade
);
