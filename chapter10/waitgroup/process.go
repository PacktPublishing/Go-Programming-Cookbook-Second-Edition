package waitgroup

import (
	"log"
	"sync"
	"time"
)

// Crawl collects responses from a list of urls
// that are passed in. It waits for all requests
// to complete before returning.
func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Printf("starting crawling")
	wg := &sync.WaitGroup{}

	var resps []int
	cerr := &CrawlError{}
	for _, v := range sites {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			resp, err := GetURL(v)
			if err != nil {
				cerr.Add(err)
				return
			}
			resps = append(resps, resp.StatusCode)
		}(v)
	}
	wg.Wait()
	if cerr.Valid() {
		return resps, cerr
	}
	log.Printf("completed crawling in %s", time.Since(start))
	return resps, nil
}
