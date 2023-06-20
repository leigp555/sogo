package test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup //等待组必须是同一个不然无法同时在一个等待组里面

// var mu *sync.Mutex =new(sync.Mutex)   // mu protects    //指针必须先初始化
var mu sync.Mutex // mu protects

var a int = 10

func test(num int) {
	fmt.Println(num)
	defer wg.Done() //完成一个协程往等待组中减一个协程
}

func add() {

	mu.Lock()
	defer mu.Unlock()
	a += 2
	wg.Done()
}
func sub() {
	mu.Lock()
	defer mu.Unlock()
	a -= 1
	wg.Done()
}

func TestMu(t *testing.T) {
	for i := 0; i < 10; i++ {
		go test(i)
		wg.Add(1) //等待组加一个协程
	}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		add()
		sub()
	}
	fmt.Println(a)
	wg.Wait() //等待所有子协程结束

}
