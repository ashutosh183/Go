package main

import (
	"fmt"
	"sync"
)


//Mutex is a kind of lock
//Only one entity can be mainpulated at the time

var mutex sync.RWMutex
var wg sync.WaitGroup
var counter int = 0
func main(){

	for i:=0; i<10; i++{
		wg.Add(2)
		mutex.Lock()
		go increment()
		mutex.RLock()
		go sayHello()
	}
	wg.Wait()
}

func sayHello(){
	fmt.Println("Called by goroutine", counter)
	mutex.RUnlock()
	wg.Done()
}

func increment(){
	counter+=1;
	mutex.Unlock()
	wg.Done()
}