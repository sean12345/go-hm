package main

import (
	//"github.com/gen-go/auction"
	"github.com/gen-go/bail"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
	"log"
)

func thriftClient() {
	startTime := currentTimeMillis()
	url := "http://192.168.59.88:80/Bail/run"

	transport, err := thrift.NewTHttpPostClient(url)
	if err != nil {
		log.Fatal("Error thrift 18 opening http failed: %v", err)
	}
	
	var protocol thrift.TProtocol = thrift.NewTBinaryProtocolTransport(transport)
	protocol = thrift.NewTMultiplexedProtocol(protocol, "BailService")
	client := bail.NewBailServiceClientProtocol(transport, protocol, protocol)
	if err := transport.Open(); err != nil {
		log.Fatal("Error thrift 25 opening transport failed: %v", err)
	}
	defer client.Transport.Close()

	var dealerID int64 = 5
	bk, err := client.GetDealerBail(dealerID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bk)

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}