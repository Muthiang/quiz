package main
import "sync"

type wordMap map[byte][]string

func (wm wordMap) loadWord(w string) {
	var l byte = w[0]
	arr := wm[l]
	wm[l] = append(arr, w)
}

func (wm wordMap) getWordArrayFor(l byte) (words []string) {
	return wm[l]
}

func (wm wordMap) getWordForProcessing(out chan<- string) {
	for _, wa := range wm {
		for _, w := range wa {
			out <- w
		}
	}
	close(out)
}

func (wm wordMap) processWord(w string, out chan<- string, wgrp *sync.WaitGroup) {
	var lh letterHash
	lh.load(w)
	for i,siz:=0,len(w); i<siz; i++ {
		lh.lookupWords(wm[w[i]], out)
	}
	wgrp.Done()
}
