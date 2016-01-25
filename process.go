package main
import "sync"

var concurrency int = 20

func process(wm wordMap, wordCount int, out chan<- string) {
	wchan := make(chan string, concurrency)
	var wgrp sync.WaitGroup
	wgrp.Add(wordCount)
	go wm.getWordForProcessing(wchan)
	for i:=0; i<concurrency; i++ {
		go func() {
			for {
				w, ok := <-wchan
				if !ok {
					return
				}
				wm.processWord(w, out, &wgrp)
			}
		}()
	}
	wgrp.Wait()
	close(out)
}

