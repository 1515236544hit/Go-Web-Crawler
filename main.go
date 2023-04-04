package main

import (
	"fmt"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type parser interface{}

type DefaultParser struct{}

//userAgents

func randomUserAgent() {

}

func extractSiteMapUrls(startURL string) []string {

	worklist := make(chan []string)
	toCrawl := []string{}

	go func() {
		worklist <- []string{startURL}
	}()

	list := <-worklist

	for _, link := range list{
		go func (link string)  {
			response, err := makeRequest(link)
			if err != nil {
				fmt.Printf("Error retriving from URl %s :",link)
			}
			urls ,_  := extractURls(response)

			if err != nil{
				fmt.Printf("Error extractig from document response URL to : %s",link)
			}

			sitemapFiles, pages := isSitemapURL(urls)
			if sitemapFiles != nil{
				worklist <- sitemapFiles
			}

			for _, page := range pages{
				toCrawl = append(toCrawl, page)
			}
		}(link)
	}
	return toCrawl
}

func makeRequest() {

}

func scrapeUrls() {

}

func scrapePage() {

}

fun extractURls(){

}

func isSitemapURL(){

}

func crawlPage() {

}

func getSEOData() {

}

func scrapeSiteMap(url string) []SeoData {

	result := extractSiteMapurls(url)
	output := scrapeUrls(result)
	return output
}

func main() {

	defaultParser := DefaultParser{}
	result := scrapeSiteMap("")

	for _, res := range result {
		fmt.Println(res)
	}
}
