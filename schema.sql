create database anaia_beauty;
use anaia_beauty;

create table ROLES (
    id_role int not null auto_increment,
    role_name varchar(50) not null,
    primary key (id_role)
);

create table USERS (
    id_user int not null auto_increment,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    role_id int not null,
    primary key (id_user),
    foreign key (role_id) references ROLES(id_role)
);

create table COURSES (
    id_course int not null auto_increment,
    title varchar(200) not null,
    description text not null,
    price decimal(10,2) not null,
    image_url varchar(255) not null,
    creation_date timestamp,
    user_id int not null,
    primary key (id_course),
    foreign key (user_id) references USERS(id_user)
);

create table CLASSES (
    id_class int not null auto_increment,
    title varchar(200) not null,
    description text not null,
    video_url varchar(255) not null,
    duration int not null,
    course_id int not null,
    primary key (id_class),
    foreign key (course_id) references COURSES(id_course)
);

create table PURCHASES (
    id_purchase int not null auto_increment,
    course_id int not null,
    purchase_date timestamp,
    amount decimal(10,2) not null,
    primary key (id_purchase),
    user_id int not null,
    foreign key (user_id) references USERS(id_user),
    foreign key (course_id) references COURSES(id_course)
);
