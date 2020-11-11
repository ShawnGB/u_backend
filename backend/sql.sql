create table users (
    UserId int primary key auto_increment not null,
    FirstName varchar(50) not null,
    LastName varchar(50) not null,
    Instructor bool
);
insert into users(FirstName, LastName, Instructor)
values ("Gordon", "Shawn", false);