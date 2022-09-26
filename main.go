package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var BaseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchType=search&searchword=python"

type extractedJob struct {
	id       string
	company  string
	title    string
	location string
}

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 1; i < totalPages; i++ {
		extractedjobs := getPage(i)
		jobs = append(jobs, extractedjobs...)
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := BaseURL + "&recruitPage=" + strconv.Itoa(page)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	company := cleanString(card.Find(".corp_name").Text())
	title := cleanString(card.Find(".job_tit>a").Text())
	location := cleanString(card.Find(".job_condition>span>a").Text())
	return extractedJob{
		id:       id,
		company:  company,
		title:    title,
		location: location}
}

//배열의 스페이스 없이 오직 텍스트만 갖고싶을때
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
