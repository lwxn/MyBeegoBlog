package models

import (
	"strings"
)

func GetMapOfTag()(map[string]int){
	tagList := QueryArticleWithParam("tags")
	tagMap := make(map[string]int)
	for _,tag := range tagList{
		tag_list := strings.Split(tag,"&")
		for _,t := range tag_list{
			tagMap[t]++
		}
	}
	return tagMap
}

func GetArticleByTags(tagValue string)([]Article){
	sql := "where tags like '%&"+ tagValue + "&%'" +  //中间
		" or tags like '" + tagValue + "&%'" +  //前面
		" or tags like '%&" + tagValue + "'"  +  //后面
		" or tags = '" + tagValue + "'"              //刚刚好
	articles,_ := QueryArticleWithCon(sql)
	return articles
}
