package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func main() {
	var dictionary string // Path to dictionary file
	var quiet bool        // If true, only show HTTP 200 responses
	var target string     // Target URL to enumerate
	var workers int       // Number of concurrent workers

	// Define command-line flags
	flag.StringVar(&dictionary, "d", "", "Path to dictionary file")
	flag.BoolVar(&quiet, "q", false, "If true, only show HTTP 200 responses")
	flag.StringVar(&target, "t", "", "Target URL to enumerate")
	flag.IntVar(&workers, "w", 1, "Number of workers to run")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of mygb:\n")
		flag.PrintDefaults()
	}
	

	flag.Parse()

	

	// Validate dictionary path
	if dictionary == "" {
		log.Fatal("Error: The dictionary path (-d) is required.")
	}

	_, err := os.Stat(dictionary)
	if os.IsNotExist(err) {
		log.Fatalf("Error: %s does not exist or is not a valid file.\n", dictionary)
	}

	// Validate target URL
	if target == "" {
		log.Fatal("Error: The target (-t) is required.")
	}

	parsedURL, err := url.Parse(target)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Fatal("Error: Invalid target URL.")
	}

	// Display configuration details if not in quiet mode
	if !quiet {
		fmt.Printf("Starting my gb\n\n")
		fmt.Printf("---\n")
		fmt.Printf("Target: %s\n", parsedURL)
		fmt.Printf("List: %s\n", dictionary)
		fmt.Printf("Workers: %d\n", workers)
		fmt.Printf("---\n\n")
		fmt.Printf("Starting scan...\n")
	}

	// Open the dictionary file
	file, err := os.Open(dictionary)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Channel for tasks
	tasks := make(chan string, workers)
	var wg sync.WaitGroup // Synchronization of goroutines

	// Launch worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(parsedURL.String(), tasks, quiet, &wg)
	}

	// Read dictionary file line by line and send paths to workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- scanner.Text()
	}

	// Close the channel to signal no more tasks
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Worker function to process URLs and send HTTP requests
func worker(baseURL string, tasks <-chan string, quiet bool, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement WaitGroup counter when function ends

	for path := range tasks {
		// Construct the full URL
		url := fmt.Sprintf("%s/%s", baseURL, path)
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Error fetching URL %s: %v", url, err)
			continue
		}
		defer response.Body.Close()

		// Print results based on -q flag
		if !quiet || response.StatusCode == 200 {
			fmt.Printf("/%-30s\t%v\n", path, response.StatusCode)
		}
	}
}
