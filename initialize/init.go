package initialize

import (
	"Irispom/cron"
	"fmt"
	"sync"
	"time"
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

//channel案例10鸡蛋100个抢（channel实现资源争抢案例）
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


//生产消息
func produceMsg(){
	//生产mq,每秒去生产一个消息
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i<=10;i++  {
		go cron.InitProduce(i, &wg)
		time.Sleep(1*time.Second)
	}
	wg.Wait()
}

//消费消息
func consumeMsg(){
	//消费mq
	go cron.InitConsumer()
}