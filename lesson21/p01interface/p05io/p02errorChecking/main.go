package main

import (
	"strings"
	"io"
	"os"
	"bytes"
	"net/http"
	"fmt"
)

func main() {

	msg := "This is sample of string value"

	rdr1 := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr1)
	if err != nil{
		fmt.Println(err)
	}

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil{
		fmt.Println(err)
	}

	resp, err := http.Get("https://dou.ua/calendar/20240/")
	defer resp.Body.Close()
	if err != nil{
		fmt.Println(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil{
		fmt.Println(err)
	}
}
