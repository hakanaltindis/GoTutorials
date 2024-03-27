package main

import (
	"fmt"
	"net/http"
	"time"
)

// version 1 - synchronous
// func main() {
// 	startTime := time.Now()
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		checkUrl(url)
	}

	fmt.Println("Execution time:", time.Since(startTime))
}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}

/*******************************************************************/
// version 2 - go
// func main() {
// 	startTime := time.Now()
// 	urls := []string{
// 		"https://www.easyjet.com/",
// 		"https://www.skyscanner.de/",
// 		"https://www.ryanair.com",
// 		"https://wizzair.com/",
// 		"https://www.swiss.com/",
// 	}
// 	for _, url := range urls {
// 		go checkUrl(url)
// 	}

// 	fmt.Println("Execution time:", time.Since(startTime))
// }

// func checkUrl(url string) {
// 	_, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(url, "is down !!!")
// 		return
// 	}
// 	fmt.Println(url, "is up and running.")
// }

/*******************************************************************/
// version 3 - time.Sleep()
// func main() {
// 	startTime := time.Now()
// 	urls := []string{
// 		"https://www.easyjet.com/",
// 		"https://www.skyscanner.de/",
// 		"https://www.ryanair.com",
// 		"https://wizzair.com/",
// 		"https://www.swiss.com/",
// 	}
// 	for _, url := range urls {
// 		go checkUrl(url)
// 	}

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Execution time:", time.Since(startTime))
// }

// func checkUrl(url string) {
// 	_, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(url, "is down !!!")
// 		return
// 	}
// 	fmt.Println(url, "is up and running.")
// }

/*******************************************************************/
// version 4 - Wait Groups
func main() {
	startTime := time.Now()
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			checkUrl(url)
		}(u)
	}

	wg.Wait()

	fmt.Println("Execution time:", time.Since(startTime))
}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}

/*******************************************************************/
// version 5 - Channels
func main() {
	startTime := time.Now()
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	c := make(chan urlStatus, 3)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		urlStats := <-c
		if urlStats.status {
			fmt.Println(urlStats.url, "is up and running.")
		} else {
			fmt.Println(urlStats.url, "is down !!!")
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}

func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- urlStatus{url, false}

	} else {
		c <- urlStatus{url, true}
	}
}

type urlStatus struct {
	url    string
	status bool
}
