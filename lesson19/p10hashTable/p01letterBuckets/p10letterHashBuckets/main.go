package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://drive.google.com/open?id=1__mXLWlRC1Y6WJqte_LSN5q4yM4XNOan")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 2000)

	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

	fmt.Println(buckets[65:123])

}

func hashBucket(word string) int {

	return int(word[0])
}
