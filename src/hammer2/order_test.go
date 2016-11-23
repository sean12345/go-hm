// AuctionHammer project main.go
package hammer2

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_start(t *testing.T) {

	var orders []*Order
	//拍0定时单
	order := NewOrder()
	order.order_id = 1
	order.order_no = "P1609102326"
	order.scene_id = 3
	order.start_time, _ = strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	order.est_second = 5
	order.real_second = 0

	go order.Start()
	orders = append(orders, order)

	//拍1定时单
	order1 := NewOrder()
	order1.order_id = 2
	order1.order_no = "P1609102327"
	order1.scene_id = 3
	order1.start_time, _ = strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	order1.est_second = 5
	order1.real_second = 0

	go order1.Start()
	orders = append(orders, order1)

	for i := range orders {
		x := orders[i].WaitEnd()
		fmt.Println(x)
	}
	fmt.Println("secne----complete")
}
