package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
)

func SwitchMarkdownToHtml(content string)template.HTML{
	makedown := blackfriday.MarkdownCommon([]byte(content))

	//获取DOM
	doc,_ := goquery.NewDocumentFromReader(bytes.NewReader(makedown))
	doc.Find("code").Each(func(i int,selection *goquery.Selection){
			light,_ := syntaxhighlight.AsHTML([]byte(selection.Text()))
			selection.SetHtml(string(light))
			fmt.Println(selection.Html())
	})
	htmlString,_ := doc.Html()
	return template.HTML(htmlString)
}

func MD5(str string) string{
	md5str := fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}
