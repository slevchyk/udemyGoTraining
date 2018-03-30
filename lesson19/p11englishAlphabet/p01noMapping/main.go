package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatal(err)
	}

	byteSlice, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	s := string(byteSlice)

	fmt.Println(s)
}