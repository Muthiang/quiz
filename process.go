package main
import "sync"

func process(wm wordMap, wordCount int, out chan<- string) {
	wchan := make(chan string, 20)
	var wgrp sync.WaitGroup
	wgrp.Add(wordCount)
	go wm.getWordForProcessing(wchan)
	go func() {
		for i:=0; i<wordCount; i++ {
			select {
				case w := <-wchan:
					go wm.processWord(w, out, &wgrp)
			}
		}
	}()
	wgrp.Wait()
	close(out)
}

