package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"log"
)

func main() {

	res, err := http.Get("https://drive.google.com/open?id=1__mXLWlRC1Y6WJqte_LSN5q4yM4XNOan")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

}
