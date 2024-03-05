package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type MapSample struct {
	sync.RWMutex
	items map[int]string
}

func (m *MapSample) Set(key int, value string) {
	m.Lock()
	defer m.Unlock()
	m.items[key] = value
}
func main() {
	map_ := MapSample{items: make(map[int]string)}
	for i := 0; i < 100; i++ {
		go func(i int) {
			map_.Set(i, fmt.Sprintf("value - %v", i))
		}(i)
	}
	fmt.Println(map_.items)
}
