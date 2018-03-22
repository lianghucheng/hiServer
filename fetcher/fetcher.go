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

func determineDecoding(r *bufio.Reader)encoding.Encoding{
	bytes,err:=r.Peek(1024)
	if err !=nil{
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}