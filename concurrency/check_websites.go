package concurrency

// WebsiteChecker checks a url, return a bool
type WebsiteChecker func(string) bool


// CheckWebsites takes a WebsiteChecker and a slice of urls and returns a map
// of urls to the result of checking each url with WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        results[url] = wc(url)
    }

    return results
}


