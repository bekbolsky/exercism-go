package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	m := make(FreqMap)
	res := make([]FreqMap, len(l))
	for i, s := range l {
		wg.Add(1)
		go func(i int, s string) {
			res[i] = Frequency(s)
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	for _, r := range res {
		for k, v := range r {
			m[k] += v
		}
	}
	return m
}

/* Realisation of ConcurrentFrequency using channels
m := FreqMap{}
ch := make(chan FreqMap)
for _, s := range l {
	go func(s string) {
		ch <- Frequency(s)
	}(s)
}
for range l {
	for k, v := range <-ch {
		m[k] += v
	}
}
return m
*/
