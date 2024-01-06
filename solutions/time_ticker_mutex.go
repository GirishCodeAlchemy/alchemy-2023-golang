package solutions

import (
	"fmt"
	"sync"
	"time"
)

func tickerFunction() {
	// Ticker for periodic tasks
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()

	// Timer for a one-time delayed task
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")
}

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Counter_function() {
	counter := Counter{}

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	time.Sleep(time.Second) // Allow time for goroutines to finish
	fmt.Println("Counter value:", counter.GetValue())
}

type Data struct {
	value int
	mu    sync.RWMutex
}

func (d *Data) Read() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.value
}

func (d *Data) Write(newValue int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.value = newValue
}

func init() {
	data := Data{}

	go func() {
		for {
			fmt.Println("Read:", data.Read())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Write:", i)
			data.Write(i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
