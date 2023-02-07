package Model

//评论是用户对某本书的评论
//需要专门建一张表来存储评论

type Comment struct {
	Id          int    //书评id，用于为每个评论唯一标识
	ParentId    int    //父评论id，用于嵌套评论
	BookId      int    //对某本评论的书籍id
	PublishTime string //发布时间
	Content     string //书评内容
	UserId      int    //发表评论的用户id
	Avatar      string //用户头像
	NickName    string //用户昵称
	//此条在被点赞的时候需要自增一次
	PraiseCount int //书评的总点赞数
	//此下为登录后才显示的内容，不登录下面两条默认为false，已登录就记录为该用户实际是否关注
	//额外建一张表，用户评论点赞表
	IsPraised bool //是否被点赞，当前用户是否给该评论点赞
	//新建表，用户与用户关系表
	IsFocus bool //当前用户是否关注发评论的用户
}

type CommentPraise struct {
	Id            int //唯一标识id
	BookId        int
	CommentId     int //书籍id
	PraiseUserId  int //点赞的用户的id
	PraisedUserId int //被点赞的用户id
}

type UserFocus struct {
	Id            int
	FocusUserId   int
	FocusedUserId int
}

//建表语句

/*
评论
create table `comment`(
	`id` bigint unsigned auto_increment primary key,
	`parent_id` bigint default 0,
	`book_id` bigint not null,
	`publish_time` varchar(20) not null,
	`content` varchar(500) not null,
	`user_id` bigint not null,
	`avatar` varchar(100) not null default "null",
	`nickname` varchar(20) not null default "null",
	`praise_count` int default 0
);
*/

/*
评论点赞表
create table `comment_praise`(
	`id` bigint unsigned auto_increment primary key,
	`book_id` bigint not null,
	`comment_id` bigint not null,
	`praise_user_id` bigint not null,
	`praised_user_id` bigint not null
);

*/

/*
create table `user_focus`(
	`id` bigint unsigned auto_increment primary key,
	`focus_user_id` bigint,
	`focused_user_id` bigint
);
*/
