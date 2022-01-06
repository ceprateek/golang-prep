package queue

import (
	"container/list"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func findFirstNonRepeatingChar(in chan rune) {
	dll := list.New()
	currentElements := make([]*list.Element, 256)
	repeated := make([]bool, 256)
	for c := range in {
		if !repeated[c] {
			if currentElements[c] == nil {
				dll.PushBack(c)
				currentElements[c] = dll.Back()
			} else {
				existing := currentElements[c]
				repeated[c] = true
				dll.Remove(existing)
			}
		}
		if dll.Len()>0 {
			fmt.Printf("first non repeating char: %v\n", string(dll.Front().Value.(int32)))
		}else {
			fmt.Println("no non repeating char")
		}
	}
	wg.Done()
}

func PlayFindFirstNonRepeatingChar() {
	input := "geeksforgeeks"
	out := make(chan rune)
	runes := []rune(input)
	wg.Add(1)
	go findFirstNonRepeatingChar(out)
	for _, c := range runes {
		out <- c
	}
	close(out)
	wg.Wait()
}
