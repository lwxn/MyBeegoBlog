package utils

func CreateTableWithArticle(){
	sql := `create table if not exists article(
				id int(4) primary key auto_increment not null,
				title varchar(64),
				author varchar(20),
				tags varchar(30),
				short varchar(255),
				content longtext,
				createtime int(10)
			);`
	ModifyDB(sql)
}

func CreateTableWithAlbum(){
	sql := `create table if not exists album(
			id int(4) primary key auto_increment,
			filepath varchar(255),
			filename varchar(64),
			status int(4),
			createtime int(10)
		);`
	ModifyDB(sql)
}
