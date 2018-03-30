package main

import (
	"fmt"
	"log"
	"net/http"
	"bufio"
	"os"
	"crypto/sha1"
	"encoding/base64"
)

func main() {

	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatal(err)
	}

	words := make(map[string][]string)

	sc := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		word := sc.Text()
		hash := hashSHA1(word)
		words[hash] = append(words[hash], word)
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	for key, val := range words {
		fmt.Println(key, val)
	}
}

func hashSHA1(word string) string {

	bs := make([]byte, 1)
	bs = append(bs, word[0])

	hasher := sha1.New()
	hasher.Write(bs)

	hashByteSlice := hasher.Sum(nil)

	hash := base64.URLEncoding.EncodeToString(hashByteSlice)
	//hash := hex.EncodeToString(hashByteSlice)

	return hash
}