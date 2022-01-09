package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct{
	rwMu sync.RWMutex
	mu sync.Mutex
	c map[string]int
}

func(c *Counter) countMe() map[string]int {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	return c.c
}

func(c *Counter) countAgain() map[string]int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.c
}

func(c *Counter) Inc(key string){
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func(c *Counter) Value(key string)int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main(){
	//package sync

	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(3 *time.Second)
	fmt.Println(c.Value(key))
}
