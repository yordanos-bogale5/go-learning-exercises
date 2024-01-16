package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	m   FreqMap
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key rune) {
	c.mux.Lock()
	c.m[key]++
	c.mux.Unlock()
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
// Pre safety
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in an array of given texts and returns this
// data as a FreqMap
func ConcurrentFrequency(t []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(t))
	c := SafeCounter{m: make(map[rune]int)}
	for _, text := range t {
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				c.Inc(r)
			}
		}(text) // function must be called with text passed in (I guess?)
	}

	wg.Wait()
	return c.m
}
