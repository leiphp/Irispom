/*
	处理定时任务，包括数据库，Redis，Mq，日志文件等
*/
package cron

import (
	"Irispom/services"
	"sync"
)

//处理MQ生产
func InitProduce(num int, wg *sync.WaitGroup) {
	//发送消息
	send := services.NewSendService()
	send.SendMsg(num, wg)
}

//处理MQ消费
func InitConsumer() {
	//接收消息
	receive := services.NewReceiveService()
	receive.ReceiveMsg()
}


