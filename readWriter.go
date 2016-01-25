package main
import (
	"bufio"
	"os"
	"fmt"
)

func readFile(wm wordMap, file string) (int, error) {
	countLines := 0
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wm.loadWord(scanner.Text())
		countLines++
	}
	return countLines, nil
}

func writeResult(out chan string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	for w := range out {
		fmt.Fprintln(f, w)
	}
	return nil
}
