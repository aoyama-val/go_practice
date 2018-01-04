package main

import (
	"fmt"
	"sync"
)

type (
	Dispatcher struct {
		queue       chan interface{}
		wg          sync.WaitGroup
		workerCount int
	}
)

func NewDispatcher(workerCount int) *Dispatcher {
	maxQueues := 10000
	d := &Dispatcher{
		queue:       make(chan interface{}, maxQueues),
		workerCount: workerCount,
	}
	return d
}

func (d *Dispatcher) Add(v interface{}) {
	d.queue <- v
}

func (d *Dispatcher) Start() {
	d.wg.Add(d.workerCount)
	for i := 0; i < d.workerCount; i++ {
		go func() {
			defer d.wg.Done()
			for v := range d.queue {
				if str, ok := v.(string); ok {
					get(str)
				}
			}
		}()
	}
}

func (d *Dispatcher) Wait() {
	close(d.queue)
	d.wg.Wait()
}

func main() {
	setMaxConnection(8)

	d := NewDispatcher(8)
	lines := readLines("url.txt")
	for _, url := range lines {
		fmt.Println(url)
		d.Add(url)
		//get(line)
	}
	d.Start()
	d.Wait()
}
