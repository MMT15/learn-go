package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			value := n
			time.Sleep(time.Duration(n) * 50 * time.Millisecond) 
			m.Store(key, value)
			fmt.Printf("Stored: %s => %d\n", key, value)
		}(i)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			time.Sleep(time.Duration(n) * 70 * time.Millisecond) 
			val, ok := m.Load(key)
			if ok {
				fmt.Printf("Loaded: %s => %v\n", key, val)
			} else {
				fmt.Printf("Key %s not found. Using LoadOrStore.\n", key)
				actual, _ := m.LoadOrStore(key, -1)
				fmt.Printf("LoadOrStore result: %s => %v\n", key, actual)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Final map contents:")
	m.Range(func(key, value any) bool {
		fmt.Printf("%v => %v\n", key, value)
		return true
	})
}
