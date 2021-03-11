package models

import (
	"BeegoDemo2/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type Article struct {
	Id int
	Title string
	Tags string
	Short string
	Content string
	Author string
	Createtime int64
}


//--------------------------------------插入文章----------------------

func insertArticle(article Article)(int64,error){
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) "+
		"values(?,?,?,?,?,?)",article.Title,article.Tags,article.Short,article.Content,article.Author,article.Createtime)
}

//-------------------------------------查询文章-----------------------
func FindArticleWithPage(page int)([]Article,error){
	//从配置文章中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page,num)
}

func FindArticleWithId(Id int)(Article){
	sql := fmt.Sprintf("where Id=%d",Id)
	artlist,err := QueryArticleWithCon(sql)
	var article Article

	if err != nil{
		fmt.Println("查询文章ID出错！！")
		return article
	}
	if artlist != nil{
		return artlist[0]
	}
	fmt.Println("查不到文章的ID！")
	return article
}


//-----------------------------------获取多少条数据-------------------
func QueryArticleWithPage(page int,num int)([]Article,error){
	sql := fmt.Sprintf("limit %d,%d",page*num,num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string)([]Article,error){
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows,err := utils.QueryDB(sql)
	if err != nil{
		return nil,err
	}
	var artList []Article
	for rows.Next(){
		var(
			id int
			title string
			tags string
			short string
			content string
			author string
			createtime int64
		)
		rows.Scan(&id,&title,&tags,&short,&content,&author,&createtime)
		art := Article{id,title,tags,short,content,author,createtime}
		artList = append(artList, art)
	}
	return artList,nil
}

//搜对应的字段
func QueryArticleWithParam(param string)([]string){
	rows,_:= utils.QueryDB(fmt.Sprintf("select %s from article ",param))
	fmt.Println("begin!")

	var tagList []string
	for rows.Next(){
		tag := ""
		rows.Scan(&tag)
		tagList = append(tagList, tag)
	}
	fmt.Println("____________________________")
	return tagList
}




//-----------------------------------翻页----------------------------------
var articleRowsNum = 0

//查询文章总数
func QueryArticleRowNum()int{
	sql := "select count(id) from article"
	result := utils.QueryRowDB(sql)

	num := 0
	result.Scan(&num)
	return num
}

//改进每次都要查询文章总数的缺点
func GetArticleRowNum()int{
	if articleRowsNum == 0{
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

func SetArticleRowsNum(){
	articleRowsNum = QueryArticleRowNum()
}

func AddArticle(article Article) (int64, error){
	count,err := insertArticle(article)
	SetArticleRowsNum()
	return count,err
}

//------------------------------------Update articles-----------------------------------
func UpdateArticleWithNew(article Article)(int64,error){
	fmt.Println(article.Createtime)
	return utils.ModifyDB("update article set title = ?,tags = ?, short = ?, createtime = ?," +
		"content = ? where id = ?",
		article.Title,article.Tags,article.Short,article.Createtime,article.Content,article.Id)

}


//------------------------------------delete articles----------------------------------------
func DeleteArticleWithId(id int)(int64,error){
	return utils.ModifyDB("delete from article where id = ?",id)
}