package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mem := make(map[string]struct{})
	var s struct{}

	for scanner.Scan() {
		title := scanner.Text()
		scanner.Scan()
		url := scanner.Text()
		scanner.Scan()
		snip := scanner.Text()

		if _, ok := mem[snip]; !ok {
			fmt.Println(title)
			fmt.Println(url)
			fmt.Println(snip)

			mem[snip] = s

		}

	}

}
