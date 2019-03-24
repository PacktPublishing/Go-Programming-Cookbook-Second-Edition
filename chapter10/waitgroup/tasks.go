package waitgroup

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetURL gets a url, and logs the time it took
func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	log.Printf("completed getting %s in %s", url, time.Since(start))
	return resp, err
}

// CrawlError is our custom error type
// for aggregating errors
type CrawlError struct {
	Errors []string
}

// Add adds another error
func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

// Error implements the error interface
func (c *CrawlError) Error() string {
	return fmt.Sprintf("All Errors: %s", strings.Join(c.Errors, ","))
}

// Valid can be used to determine if
// we should return this
func (c *CrawlError) Valid() bool {
	return len(c.Errors) != 0
}
