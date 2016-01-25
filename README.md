# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

Instructions:
copy all the files into a directory quiz.
compile using "go build" to get quiz in linux or quiz.exe in windows
wherever you run it from, have the file word.list as the main program quizmain.go uses this file to open it
the program generates an output file result.list

The design is as under:
main function reads the input file and sets up a map of value words ([]string) against first letter of the word which is the key.
The main program sets up and output string channel (chan string)
The main program calls a go routine process with the map, number of lines and the output channel.
The process function sets up a buffered (20 or equal to cores in the machine) channel for reading strings from map in a go routine of map.getWordForPorcessing which at the most supplies at a time 20 words for processing in parallel.
The process function uses a sync.WaitGroup for channel synchronization so that waiting is done, it can close the output channel.
Each word is processed in a go routine within the map called processWord.
The processWord function sets up a letterHash structure and its functions for efficient mapping using a has map against each letters of the work with their locations.
At a time 20 words are processed. At this point the map is read only and therefore concurrent processes can read.
The output from each processing is send to the output channel which reads the output one word match at a time and outputs the result.
The program uses optimum structures for space and time processing. Concurrency is used in word processing using go routines in parallel. The while program is designed and built for exploiting the currency features of the go language.
