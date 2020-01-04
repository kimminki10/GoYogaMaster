/* privilege */
/* set root password */
update mysql.user set authentication_string=password('myAldtmf956042*') where user='root';
create user 
flush privileges;

/* create database user */
create database yogamaster default character set utf8;
grant all privileges on yogamaster.* to yomauser@'%' identified by 'myAldtmf956042*';


/* create table */
/* user table 12 rows */
create table user(email varchar(1024) primary key) ENGINE=INNODB;
alter table user
    add column password_hash varchar(256) not null,
    add column u_f_name varchar(32) not null,
    add column u_l_name varchar(32),
    add column mobile varchar(12),

    add column height float,
    add column weight float,
    add column age int,
    add column gender varchar(8), 

    add column u_type varchar(16),
    add column last_used datetime,

    add column g_tkn varchar(1024),
    add column y_tkn varchar(1024);
    /* os */
    /* device type */
    /* contentslist (history) */
    /* reactions (comments, likes) */
    /* analysis */

/* pose table 7 rows */
create table pose(
    _id int primary key AUTO_INCREMENT
)ENGINE=INNODB;
alter table pose
    add column p_name varchar(1024),
    add column p_description varchar(1024), 
    add column p_thumbnail_url varchar(4096),
    add column p_video_url varchar(4096),
    add column p_json_url varchar(4096),
    add column p_category varchar(1024);
    /* TODO: effect */
    /* TODO: veiwed user */

/* poselist table */
create table poselist(
    _id int primary key AUTO_INCREMENT
)ENGINE=INNODB;
alter table poselist
    add column pose_id int,
    add constraint pose_poselist
    FOREIGN KEY(pose_id) REFERENCES pose(_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

/* contents table */
create table contents(
    _id int primary key AUTO_INCREMENT
)ENGINE=INNODB;

alter table contents
    add column email varchar(1024),
    add constraint contents_user
    FOREIGN KEY(email) REFERENCES user(email)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

    add column c_name          varchar(1024),
    add column c_description   varchar(4096),
    add column durations       float,
    add column category        varchar(256),
    add column c_thumbnail_url varchar(4096),
    add column c_view_num      int,
    add column c_status        varchar(16),
    
    ADD COLUMN poselist_id int,
    add constraint contents_poselist
    FOREIGN KEY(poselist_id) REFERENCES poselist(_id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;
    /* relatives 연관 강좌들 */
    /* userlist 수강생 목록 */
    /* status (ongoing, done) */
    /* reactions */

