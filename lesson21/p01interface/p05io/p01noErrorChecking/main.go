package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	msg := "This is sample of string value"

	rdr1 := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr1)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	resp, _ := http.Get("https://dou.ua/calendar/20240/")
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
