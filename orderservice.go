// AuctionHammer project main.go
package main

import (
	"fmt"
	"strconv"

	"time"
)

var sync_chan chan int //同步chan

type OrderService struct {
	//redis['start_list_'+秒值]='此秒开始的拍单列表'
	//redis['end_list_'+秒值]='此秒结束的拍单列表'
}

//构造
func NewOrder() *OrderService {

}

//每秒分发器
func distributerForSecond() {
	//每秒起动一个协程
	//每秒检测redis是否有需要开始和结束的列表
	//有,启动拍单分发器，分发处理
	go distributerForOrder()
}

//拍单处理分发器
func distributerForOrder() {
	//从开始列表POP出一条处理开始
	go orderStart()
	//从结束列表POP出一条处理结束
	go orderEnd()
}

//处理拍单开始
func orderStart() {

}

//处理拍单结束
func orderEnd() {

}
