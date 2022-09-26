package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var BaseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchType=search&searchword=python"

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := BaseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println(pageURL)
}

func getPages() int {
	// res := getHttp(BaseURL)
	pages := 0
	res, err := http.Get(BaseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages

}

/*
http.Get 이용시 403 status코드가 나옴. newrequest를 통해 request객체를 만들고 header를 추가(사이트 마다 다름)

func getHttp(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Crawler")
	client := &http.Client{} //http client를 통해서 request 실행
	res, rErr := client.Do(req)
	checkErr(rErr)
	checkCode(res)
	return res
}
*/

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}
