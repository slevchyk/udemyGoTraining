package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

func main()  {

	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")

	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
}
