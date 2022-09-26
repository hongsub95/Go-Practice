package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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
	c := make(chan []extractedJob)
	totalPages := getPages()
	for i := 1; i < totalPages; i++ {
		go getPage(i, c)
	}
	for i := 1; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}
	writeJobs(jobs)
	fmt.Println("당신이 추출한 갯수는", len(jobs), "개 입니다")
}

//defer는 함수 끝날때 "이거 하겠다"라고 선언하는 명령문
func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := BaseURL + "&recruitPage=" + strconv.Itoa(page)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c) //go 루틴을 만들고, extractjob함수에 channel인자를 보낸다
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

//channel을 통해 보낼 메세지들을 return
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	company := cleanString(card.Find(".corp_name").Text())
	title := cleanString(card.Find(".job_tit>a").Text())
	location := cleanString(card.Find(".job_condition>span>a").Text())
	c <- extractedJob{
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

// csv형태로 저장, 함수끝나는 시점에 defer를 이용하여 flush로 저장
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Company", "Title", "Location"}
	wErr := w.Write(headers)
	checkErr(wErr)
	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.company, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)

	}
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
