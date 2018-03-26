package  main

import "fmt"

func main() {

	switch "Serhii" {
	case "Jane":
		fmt.Println("Hello Jane")
	case "Juliia":
		fmt.Println("Hello Juliia")
	case "Serhii":
		fmt.Println("Hello Serhii")
		fallthrough
	case "Viktoriia":
		fmt.Println("Hello Victoriia")
	case "Andrii":
		fmt.Println("Hello Andrii")
	}
}
