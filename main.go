// AuctionHammer project main.go
package main

import (
	//"bufio"
	"fmt"
	"os"
	//	"strconv"
	"sync"
	"time"
)

var log_file = "a.txt"
var wg sync.WaitGroup

func File_put(txt string) {
	fd, _ := os.OpenFile(log_file, os.O_APPEND, 0644)
	fmt.Fprintln(fd, txt)

	defer fd.Close()

}

func a() {
	File_put("defer----")
	wg.Done()
}

//var order_chan = make(chan int)

//func order() chan int {
//	c := time.After(time.Second * 5)

//	select {
//	case <-c:
//		go a()
//	}
//	wg.Done()
//	return c
//}

func order_end(t time.Time) {
	//TODO 拍单结束处理
	File_put("order--end" + t.String())
}
func order__(wg sync.WaitGroup, t time.Time) {
	File_put("order--start" + t.String())
	c := time.After(time.Second * 5)
	select {
	case <-c:
		order_end(t)
	}
	wg.Done()
}

func secene(t time.Time) {
	var sec_wg sync.WaitGroup
	File_put("secene--" + t.String())
	sec_wg.Add(1)
	if t.Second()%5 == 0 {
		go order__(sec_wg, t)
	}

	sec_wg.Wait()
}

func main() {
	for {
		go secene(time.Now())
		time.Sleep(time.Second * 1)
	}
}
