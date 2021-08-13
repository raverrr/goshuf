package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var in string
var lines []string

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = scanner.Text()
		lines = append(lines, in)
	}
	//fmt.Println(in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	linesSlices := strings.Join(lines[:], "\n")
	fmt.Println(linesSlices)
	//fmt.Println(lines)
}
