package Model

//用于存放书的数据模型
/*
create table `book`(
	`book_id` bigint unsigned auto_increment primary key,
	`name` varchar(30) not null,
	`author` varchar(20) not null default "null",
	`score` int default 0,
	`cover` varchar(100) default "null",
	`publish_time` varchar(20) default "null",
	`link` varchar(100) default "null",
	`label` varchar(100) default "null"
);

create table `collection`(
	`id` bigint unsigned auto_increment primary key,
	`book_id` bigint not null,
	`user_id` bigint not null,
	`username` varchar(20) not null,
	`bookname` varchar(20) not null
);

create table `comments`(
	`id` bigint unsigned auto_increment primary key,
	`book_id` bigint not null,
	`commenter_id` bigint not null,
	`commenter_name` varchar(20) not null,
	`parent_id` bigint not null default 0
);
*/

type Collection struct {
	Id       int    `form:"id"`
	BookId   int    `form:"book_id"`
	BookName string `form:"book_name"`
	UserId   int    `form:"user_id"`
	UserName string `form:"username"`
}

type Book struct {
	Id     int    `form:"id"`
	Name   string `form:"name"`
	Author string `form:"author"`
	//is_star标签在用户和书籍的收藏关系表中获取，默认为false
	//还需要一张书评表
	IsStar      bool    `form:"is_star"`
	CommentNum  int     `form:"comment_num"`
	Score       string  `form:"score"`
	Cover       string  `form:"cover"`
	PublishTime string  `form:"publish_time"`
	Link        string  `form:"link"`
	Label       string  `form:"label"`
	Temperature float64 `form:"temperature"`
}
