package concurrency

type WebsiteChecker func(string) bool
type Result struct {
	url       string
	isChecked bool
}

func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	resultsChannel := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- Result{u, w(u)}
		}(url)
	}

	results := make(map[string]bool)
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.url] = r.isChecked
	}
	return results
}
