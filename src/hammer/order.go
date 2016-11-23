// AuctionHammer project main.go
package hammer

import (
	"fmt"
	"strconv"

	"time"
)

var sync_chan chan int //同步chan

type Order struct {
	order_id   int    //拍单id
	order_no   string //拍单号
	scene_id   int    //拍场id
	start_time int    //竞拍开始时间，nil此单会从场开始顺序拍
	est_second int    //预计竞拍时长 秒

	real_second int //实计竞拍时长
}

//构造
func NewOrder() *Order {
	order := new(Order)
	sync_chan = make(chan int)

	return order
}

//等待拍单结否
func (o *Order) WaitEnd() int {
	x := <-sync_chan
	return x
}

//TODO拍单开始
func (o *Order) Start() {

	o.startAuction()
}
func (o *Order) startAuction() {
	//TODO 更改拍单状态
	fmt.Println("start--" + o.order_no)
	fmt.Println("竞拍时长" + strconv.Itoa(o.getRemainSecond()))

	o.waitEndAuction()
}

//等待拍卖结束
func (o *Order) waitEndAuction() {
	dur, _ := time.ParseDuration(strconv.Itoa(o.getRemainSecond()) + "s")
	c := time.After(dur)
	select {
	case <-c:
		o.endAuction()
	}
}

//获得竞拍剩余时长
func (o *Order) getRemainSecond() int {
	return o.maxSecond() - (Int64ToInt(time.Now().Unix()) - o.start_time)
}

func (o *Order) maxSecond() int {
	max_second := o.est_second
	if o.real_second > o.est_second {
		max_second = o.real_second
	}
	return max_second
}
func Int64ToInt(n1 int64) int {
	n, err := strconv.Atoi(strconv.FormatInt(n1, 10))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}

//TODO 拍单结束
func (o *Order) endAuction() {
	//TODO 找出最后10秒是否有出价记录
	if o.real_second < (30 * o.order_id) {
		//TODO [1]有，增加竞拍时长
		//TODO延长竞拍时长
		o.real_second = o.maxSecond() + 5
		fmt.Println("defer to ---" + o.order_no + "------" + strconv.Itoa(o.real_second))
		o.waitEndAuction()
	} else {
		//TODO [2]否,结束，更改拍单状态，记录结束时间
		//TODO 结束

		sync_chan <- o.order_id
		fmt.Println("end---" + o.order_no)

		return
	}

}
