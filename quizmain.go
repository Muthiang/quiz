package main
import (
	"fmt"
	"os"
)

func main() {
	wm := make(wordMap)
	countLines, err := readFile(wm, "word.list")
	if err != nil {
		fmt.Println("could not read input")
		os.Exit(1)
	}
	out := make(chan string)
	go process(wm, countLines, out)
	if writeResult(out, "result.list") != nil {
		fmt.Println("could not write result")
		os.Exit(1)
	}
}
