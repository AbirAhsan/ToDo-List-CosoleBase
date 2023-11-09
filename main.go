package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("===============TODO LIST===============")

	scanner.Scan()
	str := scanner.Text()
	fmt.Println("Scaned Text", str)
}
