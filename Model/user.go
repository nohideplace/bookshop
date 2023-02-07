package Model

/*
 create table `user`(
     `id` bigint unsigned auto_increment primary key,
     `username` varchar(20) not null unique,
     `password` varchar(20) not null,
     `gender` varchar(10) default "null",
     `nickname` varchar(10) default "null",
     `qq` bigint default 114514,
     `birthday` varchar(20) default "null",
     `email` varchar(20) default "null",
     `avatar` varchar(100) default "null",
     `introduction` varchar(200) default "null",
     `phone` bigint default 114514
);
*/

type User struct {
	Id           int    `form:"id"`
	Username     string `form:"username"`
	Password     string `form:"password"`
	Gender       string `form:"gender"`
	Nickname     string `form:"nickname"`
	QQ           int    `form:"qq"`
	Birthday     string `form:"birthday"`
	Email        string `form:"email"`
	Avatar       string `form:"avatar"`
	Introduction string `form:"introduction"`
	Phone        int    `form:"phone"`
}
