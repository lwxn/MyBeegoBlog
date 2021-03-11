package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

//标签
type TagLink struct {
	TagName string
	TagUrl string
}
//是首页的内容
type HomeBlockParam struct {
	Id int
	Title string
	Tags []TagLink
	Short string
	Content string
	Author string
	CreateTime string
	//文章地址
	Link string

	//修改文章地址
	UpdateLink string
	DeleteLink string

	//用户是否登录
	IsLogin bool
}

//分页的对象
type HomeFooterPageCode struct {
	HasPre bool
	HasNext bool
	ShowPage string
	PreLink string
	NextLink string
}

//替换一下首页需要展示的html
func MakeHomeBlocks(articles []Article,isLogin bool)template.HTML{
	htmlHome := ""
	for _, art := range articles{
		//将article model转换首页model
		homeBlockParam := HomeBlockParam{}
		homeBlockParam.Id = art.Id
		homeBlockParam.Title = art.Title
		homeBlockParam.Tags = createTagsLinks(art.Tags)
		homeBlockParam.Short = art.Short
		homeBlockParam.Content = art.Content
		homeBlockParam.Author = art.Author
		homeBlockParam.CreateTime = strconv.FormatInt(art.Createtime,10)
		homeBlockParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeBlockParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeBlockParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeBlockParam.IsLogin = isLogin

		//解析html,插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer,homeBlockParam)
		htmlHome += buffer.String()

	}
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string)[] TagLink{
	var tagLink []TagLink
	tagsPamar := strings.Split(tags,"&")
	for _, tag := range tagsPamar{
		tagLink = append(tagLink,TagLink{tag,"/?tag="+tag})
	}
	return tagLink
}

//翻页功能
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode{
	homeFooterPageCode := HomeFooterPageCode{}
	//查询总的条数
	num := GetArticleRowNum()
	//从配置文件读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算总页数
	allPageNum_tmp := float64(num)/float64(pageRow)
	allPageNum := int(float64(num)/float64(pageRow))//int((float32(num) + 0.5))/pageRow
	if allPageNum_tmp - float64(allPageNum) > 0{
		allPageNum++
	}

	homeFooterPageCode.ShowPage = fmt.Sprintf("%d/%d",page,allPageNum)
	//上一页
	if page <= 1{
		homeFooterPageCode.HasPre = false
	}else{
		homeFooterPageCode.HasPre = true
	}
	//下一页
	if page < allPageNum{
		homeFooterPageCode.HasNext = true
	}else {
		homeFooterPageCode.HasNext = false
	}

	homeFooterPageCode.PreLink = "/?page="+strconv.Itoa(page-1)
	homeFooterPageCode.NextLink = "/?page="+strconv.Itoa(page+1)
	return homeFooterPageCode
}


