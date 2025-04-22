package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func WebsiteCheckers(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channelResult <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <- channelResult
		results[result.string] = result.bool
	}

	return results
}
