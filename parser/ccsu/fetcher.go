package ccsu

import (
	"net/http"
	"net/url"
	"net/http/cookiejar"
	"time"
	"strconv"
	"io/ioutil"
	"regexp"
	"github.com/astaxie/beego"
)


func Fectch(reqUrl string,account string,password string){
	c:=http.Client{}
	postData:=url.Values{}
	postData.Add("USERNAME",account)
	postData.Add("PASSWORD",password)

	res,err:=c.PostForm(reqUrl,postData)
	if err!=nil{

	}
	jar,err:=cookiejar.New(nil)
	if err!=nil{
	}
	parseUrl,_:=url.Parse(reqUrl)
	jar.SetCookies(parseUrl,res.Cookies())
	c.Jar=jar
	res,_=c.Get("http://jwcxxcx.ccsu.cn/jwxt/tkglAction.do?method=goListKbByXs&sql=&xnxqh=2017-2018-2&zc=&xs0101id="+account)
	body,_:=ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//beego.Debug(string(body))
	beego.Debug()
	CcsuListParser(string(body))

}

func extractString(contents []byte,re *regexp.Regexp)string {
	match:=re.FindSubmatch(contents)

	if len(match) >=2{
		return string(match[1])
	}else{
		return ""
	}
}

func StrDate() string {
	var date string
	year := Substr(time.Now().Format("2006-01-02 15:04"), 0, 4)
	month := Substr(time.Now().Format("2006-01-02 15:04"), 5, 2)
	year_num, _ := strconv.Atoi(year)
	month_num, _ := strconv.Atoi(month)
	if month_num > 8 || month_num < 2 {
		year_num++
		date = year + "-" + strconv.Itoa(year_num) + "-1"
	} else if month_num >= 2 && month_num <= 8 {
		year_num--
		date = strconv.Itoa(year_num) + "-" + year + "-2"
	}
	return date
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
