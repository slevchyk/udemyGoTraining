package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {

	const input = `It will be seen that this mere painstaking burrower and grub-worm of
	a poor devil of a Sub-Sub appears to have gone through the long
	Vaticans and street-stalls of the earth, picking up whatever random
	allusions to whales he could anyways find in any book whatsoever,
		sacred or profane. Therefore you must not, in every case at least,
	take the higgledy-piggledy whale statements, however authentic, in
	these extracts, for veritable gospel cetology. Far from it. As
	touching the ancient authors generally, as well as the poets here
	appearing, these extracts are solely valuable or entertaining, as
	affording a glancing bird's eye view of what has been promiscuously
	said, thought, fancied, and sung of Leviathan, by many nations and
	generations, including our own.`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

}
