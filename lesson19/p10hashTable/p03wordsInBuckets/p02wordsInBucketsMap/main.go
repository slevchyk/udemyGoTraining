package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://gist.githubusercontent.com/StevenClontz/4445774/raw/1722a289b665d940495645a5eaaad4da8e3ad4c7/mobydick.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	numberBuckets := 12
	buckets := make(map[int]map[string]int)

	for i := 0; i < numberBuckets; i++ {
		buckets[i] = make(map[string]int)
	}

	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, numberBuckets)
		buckets[n][word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

	fmt.Println(buckets)

	for key, val := range buckets[6] {
		fmt.Println(key, " - ", val)
	}
}

func hashBucket(word string, buckets int) int {

	var sum int

	for _, val := range word {
		sum += int(val)
	}

	bucket := sum % buckets

	return bucket
}
