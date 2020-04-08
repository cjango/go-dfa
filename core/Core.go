package core

import (
	"log"
	"sync"
)

type Result map[string][]string

type words []string

var wg sync.WaitGroup

func Check(content string, types []string) (result Result, ok bool) {
	ok = true
	checkResult := make(chan checkResult, len(types))
	result = make(Result, len(types))
	log.Println(len(types))
	wr := &sync.WaitGroup{}
	go func() {
		wr.Add(1)
		for response := range checkResult {
			if response.key != "" {
				ok = false
				result[response.key] = response.value
			}
		}
		wr.Done()
	}()

	for _, t := range types {
		wg.Add(1)
		go worker(content, t, checkResult)
	}
	wg.Wait()

	close(checkResult)

	wr.Wait()
	return
}
