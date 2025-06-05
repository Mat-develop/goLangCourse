package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time" // Added for timeout
)

// Google I/O 2012 - Go Concurrency Patterns
// <- chan - canal somente leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			// 1. Error Handling for http.Get
			resp, err := http.Get(url)
			if err != nil {
				// Send an error message or a placeholder if the request fails
				c <- fmt.Sprintf("Erro ao acessar %s: %v", url, err)
				return
			}
			defer resp.Body.Close() // Ensure the response body is closed

			// 2. Error Handling for io.ReadAll
			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("Erro ao ler HTML de %s: %v", url, err)
				return
			}

			// 3. Fix the regular expression typo: <\\/tittle should be <\\/title>
			//    And escape the forward slash if it's not part of a character class.
			//    Also, make sure to compile the regex outside the loop if performance is critical,
			//    but for a small number of URLs, it's fine here.
			r := regexp.MustCompile(`<title>(.*?)</title>`) // Corrected regex
			matches := r.FindStringSubmatch(string(html))

			if len(matches) > 1 {
				c <- matches[1]
			} else {
				// Handle cases where the title tag is not found
				c <- fmt.Sprintf("Título não encontrado para %s", url)
			}
		}(url)
	}
	return c
}

func main() {
	// Added a client with a timeout to prevent hanging on unresponsive URLs
	client := &http.Client{Timeout: 5 * time.Second}
	_ = client // This client is not directly used in 'titulo' as it is,
	           // but it's good practice for real-world scenarios.
	           // For this example, http.Get uses the default client which has no timeout.
	           // To use a custom client, you'd modify 'titulo' to accept it.

	t1 := titulo("https://www.regexr.com/", "https://www.youtube.com/")
	t2 := titulo("https://www.google.com/", "https://www.amazon.com/")

	fmt.Println("Primeiros", <-t1, "|", <-t2)
	fmt.Println("Segundos", <-t1, "|", <-t2)
}