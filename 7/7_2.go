package main

import (
	"fmt"
	"sync"
)

/*
	№ 7 (2 решение)

	Реализовать конкурентную запись данных в map.
*/

type m struct {
	mu sync.Mutex
	m  map[any]any
}

func NewMap() *m {
	return &m{
		mu: sync.Mutex{},
		m:  make(map[any]any),
	}
}

func (_m *m) Set(key any, value any) {
	_m.mu.Lock()
	defer _m.mu.Unlock()

	_m.m[key] = value
}

func main() {
	m := NewMap()

	wg := &sync.WaitGroup{}

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			m.Set(index, struct{}{})
		}(index)
	}

	wg.Wait()

	fmt.Println(m.m)
}
