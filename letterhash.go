package main

// letterHash holds a hash table by 26 lower case letters assuming words are in lower case
type letterHash struct {
    letters_ [26]*letter
    largest_ string
}

// loads a given word in the hash table with each letter location for faster search
func (lh *letterHash) load(word string) {
    for i, l := range word {
        loc := int(l) - int('a')
        lh.letters_[loc] = newLetter(i, lh.letters_[loc])
    }
    lh.largest_ = word
}

// efficient comparison functon from the start location in the largest word
func cmpStringFromLoc(larger, smaller string, loc int) bool {
    size := len(larger) - loc
    if len(smaller) > size {
        return false
    }
    for i,j:=loc,0; j<len(smaller); i,j=i+1,j+1 {
        if larger[i] != smaller[j] {
            return false;
        }
    }
    return true
}

// lookup performs a faster search from the first letter location in the hash table
func (lh *letterHash) lookup(word string) bool {
	if len(lh.largest_) <= len(word) {
		return false
	}
    loc := int(word[0]) - int('a')
    var let *letter = lh.letters_[loc]
    for let != nil {
        if cmpStringFromLoc(lh.largest_, word, let.loc_) {
            return true
        }
        let = let.next_
    }
    return false
}

// performs if words are subset words of the largest word
func (lh *letterHash) lookupWords(words []string, out chan <- string) {
    for _, s := range words {
        if lh.lookup(s) {
			out <- lh.largest_ + ": " + s
        }
    }
}
