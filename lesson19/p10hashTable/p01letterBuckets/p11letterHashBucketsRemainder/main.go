package main

import (
	"log"
	"bufio"
	"fmt"
	"os"
	"net/http"
)

func main() {

	res, err := http.Get("https://drive.google.com/open?id=1__mXLWlRC1Y6WJqte_LSN5q4yM4XNOan")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	numberBuckets := 12
	buckets := make([]int, numberBuckets)

	for scanner.Scan() {
		n := hashBucket(scanner.Text(), numberBuckets)
		buckets[n]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

	fmt.Println(buckets)

}

func hashBucket(word string, buckets int) int {

	letter := int(word[0])
	bucket := letter % buckets

	return bucket
}