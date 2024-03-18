package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type NewMap struct {
	sync.RWMutex
	mp map[int]int
}

// Функция добавления в мап
func (n *NewMap) add(key, value int) {
	n.Lock()
	defer n.Unlock()
	n.mp[key] = value
}

// Функция получение из мап
func (n *NewMap) get(key int) (int, bool) {
	n.RLock()
	defer n.RUnlock()
	value, ok := n.mp[key]
	return value, ok
}

// Функция вывода значения
func (n *NewMap) print() {
	n.RLock()
	defer n.RUnlock()
	for key, value := range n.mp {
		fmt.Printf("m[%d] = %d\n", key, value)
	}
}

func main() {
	m := NewMap{mp: make(map[int]int, 10)}

	var wg sync.WaitGroup
	wg.Add(10)

	// Цикл для создания горутин
	for i := 0; i < 10; i++ {
		go func(m *NewMap, key int, wg *sync.WaitGroup) {
			defer wg.Done()
			m.add(key, rand.Intn(100))
		}(&m, i, &wg)
	}

	wg.Wait()
	m.print()
}
