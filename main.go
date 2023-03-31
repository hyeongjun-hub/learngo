package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseURL = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=devops"
var jobCount = 40 // 사람인 페이지 하나가 가지고 있는 공고 개수

type job struct {
	id         string
	enterprise string
	title      string
	location   string
}

func main() {
	totalPages := getPages()
	c := make(chan job)
	fmt.Println("I have", totalPages, "pages")

	for i := 1; i <= totalPages; i++ {
		go getPage(i, c)
	}

	var jobs []job

	for i := 1; i <= jobCount*totalPages; i++ {
		j := <-c
		jobs = append(jobs, j)
	}

	writeJobs(jobs)
	fmt.Println("추출 완료")
}

// getGoQuery: url 의 http response 를 받는 함수
func getGoQuery(URL string) (*http.Response, *goquery.Document) {
	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	return res, doc
}

// getPages: 해당 공고가 몇개의 페이지 수를 가졌는지 세는 함수
func getPages() int {
	pages := 0
	res, doc := getGoQuery(baseURL)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		checkErr(err)
	}(res.Body)

	return pages
}

// getPage : 하나의 페이지에서 각 Card 마다 extractCard 를 호출, channel 로 전송
func getPage(page int, c chan<- job) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	res, doc := getGoQuery(pageURL)

	doc.Find(".item_recruit").Each(func(i int, s *goquery.Selection) {
		c <- extractCard(s)
	})

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		checkErr(err)
	}(res.Body)
}

// extractCard: 원하는 내용을 추출하는 함수
func extractCard(s *goquery.Selection) job {
	jobTitle := cleanString(s.Find(".job_tit").Text())
	id, _ := s.Attr("value")
	location := cleanString(s.Find(".job_condition>span>a").Text())
	enterprise := cleanString(s.Find(".corp_name>a").Text())
	return job{id: id, enterprise: enterprise, title: jobTitle, location: location}
}

// checkErr : error check
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// checkCode : http status code 를 확인하는 함수
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with Status:", res.StatusCode)
	}
}

// cleanString: string 의 white space 를 제거하는 함수
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []job) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	// Flush 함수로 csv 생성
	defer w.Flush()

	headers := []string{"Enterprise", "Title", "Location", "Link"}

	err = w.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		jobSlice := []string{job.enterprise, job.title, job.location, "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}
