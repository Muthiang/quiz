package main

// letter holds the information of suset word location in the largest word
type letter struct {
    loc_ int
    next_ *letter    
}

// letter constructor
func newLetter(loc int, next *letter) (*letter) {
    return &letter{loc_: loc, next_: next}
}
