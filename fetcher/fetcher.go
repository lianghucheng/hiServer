package fetcher

import (
	"golang.org/x/text/transform"
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"net/url"
	"github.com/astaxie/beego"
	"net/http/cookiejar"
)


func Fetch(url string) ([]byte,error){
	resp,err:=http.Get(url)
	if err !=nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("Wrong status code: %d",resp.StatusCode)
	}
	//utf8Reader:=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	bodyReader:=bufio.NewReader(resp.Body)
	e:=determineDecoding(bodyReader)
	utf8Reader:=transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func PostFetch(post_url string,account string,password string) ([]byte,error){
	c:=&http.Client{}


	postValues:=url.Values{}
	postValues.Add("USERNAME",account)
	postValues.Add("PASSWORD",password)
	res,err:=c.PostForm(post_url,postValues)
	cookie:=res.Cookies()
	beego.Debug(res.Cookies())
	if err!=nil{
		beego.Error("error")
	}
	postURL,_ := url.Parse(post_url)
	jar,_:= cookiejar.New(nil)
	jar.SetCookies(postURL,cookie) //重新新建一个cookie文件
	c.Jar=jar
	res,_=c.Get("http://www.ccsu.cn/")
	body,_:= ioutil.ReadAll(res.Body)
	return body,nil
}

func determineDecoding(r *bufio.Reader)encoding.Encoding{
	bytes,err:=r.Peek(1024)
	if err !=nil{
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}