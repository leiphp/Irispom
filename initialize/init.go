package initialize

import (
	"fmt"
	"sync"
)

//	提供系统初始化，全局变量
func Init() {
	//channel()
	eggs()
}

//channel测试方法
func channel(){
	//初始化channel
	ch := make(chan int)
	//输出channel
	go func() {
		fmt.Println(<-ch)
	}()
	//输入channel
	ch <- 5
	//关闭channel
	close(ch)
}

//channel案例10鸡蛋100个抢
func eggs(){
	//初始化eggs
	eggs := make(chan int, 10)
	//输入10个鸡蛋
	for i :=0; i<10; i++ {
		eggs <- i
	}

	//初始化wg
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			select{
			case egg:= <-eggs:
				fmt.Printf("people: %d, get egg: %d \n",num, egg)
			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}