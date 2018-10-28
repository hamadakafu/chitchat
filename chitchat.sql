-- This Sql Have To Be Done -U postgres -d chitchat

create table chitchatinfo(
    number_of_user integer not null,
    number_of_chat integer not null
);

create table chatlist(
    create_user_id integer not null,
    create_user_name integer not null,
    create_date date not null,
    chat_hash varchar(256) not null,
    chat_title varchar(256) not null,
    number_of_comment integer not null
);

create table comments(
    comment_id integer not null,
    comment_text varchar(256) not null,
    create_user_id integer not null,
    create_name varchar(256) not null,
    create_date date not null,
    chat_hash varchar(256) not null,
    primary key(comment_id)
);

create table userinfo(
    user_id int not null,
    user_name varchar(256) not null,
    user_password varchar(256) not null,
    create_date date not null,
    session_state boolean not null,
    session_id varchar(256)
);

grant all on chitchatinfo to chitchatmanager;
grant all on chatlist to chitchatmanager;
grant all on comments to chitchatmanager;
grant all on userinfo to chitchatmanager;

begin TRANSACTION;
insert into chitchatinfo values (0, 0);
commit;