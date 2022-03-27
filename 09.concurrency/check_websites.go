package concurrency

// WebsiteChecker checks a url, return a bool
type WebsiteChecker func(string) bool

type result struct {
	url_string string
	status     bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns a map
// of urls to the result of checking each url with WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url_string: u, status: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url_string] = r.status
	}

	return results
}
