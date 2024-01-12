package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	for _, url := range urls {
		result[url] = w(url)
	}
	return result
}
