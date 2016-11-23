// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package auction

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type AuctionService interface {
	// Parameters:
	//  - DealerId
	//  - OrderId
	//  - Price
	Bidding(dealerId int64, orderId int64, price float64) (r bool, err error)
	// Parameters:
	//  - DealerId
	//  - OrderId
	//  - Price
	Bid(dealerId int64, orderId int64, price float64) (r bool, err error)
	// Parameters:
	//  - Scenne
	StartAuction(scenne *Scene) (r bool, err error)
}

type AuctionServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAuctionServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AuctionServiceClient {
	return &AuctionServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAuctionServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AuctionServiceClient {
	return &AuctionServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - DealerId
//  - OrderId
//  - Price
func (p *AuctionServiceClient) Bidding(dealerId int64, orderId int64, price float64) (r bool, err error) {
	if err = p.sendBidding(dealerId, orderId, price); err != nil {
		return
	}
	return p.recvBidding()
}

func (p *AuctionServiceClient) sendBidding(dealerId int64, orderId int64, price float64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("bidding", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AuctionServiceBiddingArgs{
		DealerId: dealerId,
		OrderId:  orderId,
		Price:    price,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AuctionServiceClient) recvBidding() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "bidding" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "bidding failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "bidding failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "bidding failed: invalid message type")
		return
	}
	result := AuctionServiceBiddingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Ex != nil {
		err = result.Ex
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - DealerId
//  - OrderId
//  - Price
func (p *AuctionServiceClient) Bid(dealerId int64, orderId int64, price float64) (r bool, err error) {
	if err = p.sendBid(dealerId, orderId, price); err != nil {
		return
	}
	return p.recvBid()
}

func (p *AuctionServiceClient) sendBid(dealerId int64, orderId int64, price float64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("bid", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AuctionServiceBidArgs{
		DealerId: dealerId,
		OrderId:  orderId,
		Price:    price,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AuctionServiceClient) recvBid() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "bid" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "bid failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "bid failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "bid failed: invalid message type")
		return
	}
	result := AuctionServiceBidResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Ex != nil {
		err = result.Ex
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Scenne
func (p *AuctionServiceClient) StartAuction(scenne *Scene) (r bool, err error) {
	if err = p.sendStartAuction(scenne); err != nil {
		return
	}
	return p.recvStartAuction()
}

func (p *AuctionServiceClient) sendStartAuction(scenne *Scene) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("startAuction", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AuctionServiceStartAuctionArgs{
		Scenne: scenne,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AuctionServiceClient) recvStartAuction() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "startAuction" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "startAuction failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "startAuction failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "startAuction failed: invalid message type")
		return
	}
	result := AuctionServiceStartAuctionResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Ex != nil {
		err = result.Ex
		return
	}
	value = result.GetSuccess()
	return
}

type AuctionServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AuctionService
}

func (p *AuctionServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AuctionServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AuctionServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAuctionServiceProcessor(handler AuctionService) *AuctionServiceProcessor {

	self6 := &AuctionServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["bidding"] = &auctionServiceProcessorBidding{handler: handler}
	self6.processorMap["bid"] = &auctionServiceProcessorBid{handler: handler}
	self6.processorMap["startAuction"] = &auctionServiceProcessorStartAuction{handler: handler}
	return self6
}

func (p *AuctionServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type auctionServiceProcessorBidding struct {
	handler AuctionService
}

func (p *auctionServiceProcessorBidding) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AuctionServiceBiddingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("bidding", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AuctionServiceBiddingResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.Bidding(args.DealerId, args.OrderId, args.Price); err2 != nil {
		switch v := err2.(type) {
		case *InvalidException:
			result.Ex = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing bidding: "+err2.Error())
			oprot.WriteMessageBegin("bidding", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("bidding", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type auctionServiceProcessorBid struct {
	handler AuctionService
}

func (p *auctionServiceProcessorBid) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AuctionServiceBidArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("bid", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AuctionServiceBidResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.Bid(args.DealerId, args.OrderId, args.Price); err2 != nil {
		switch v := err2.(type) {
		case *InvalidException:
			result.Ex = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing bid: "+err2.Error())
			oprot.WriteMessageBegin("bid", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("bid", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type auctionServiceProcessorStartAuction struct {
	handler AuctionService
}

func (p *auctionServiceProcessorStartAuction) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AuctionServiceStartAuctionArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("startAuction", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AuctionServiceStartAuctionResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.StartAuction(args.Scenne); err2 != nil {
		switch v := err2.(type) {
		case *InvalidException:
			result.Ex = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing startAuction: "+err2.Error())
			oprot.WriteMessageBegin("startAuction", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("startAuction", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - DealerId
//  - OrderId
//  - Price
type AuctionServiceBiddingArgs struct {
	DealerId int64   `thrift:"dealerId,1" json:"dealerId"`
	OrderId  int64   `thrift:"orderId,2" json:"orderId"`
	Price    float64 `thrift:"price,3" json:"price"`
}

func NewAuctionServiceBiddingArgs() *AuctionServiceBiddingArgs {
	return &AuctionServiceBiddingArgs{}
}

func (p *AuctionServiceBiddingArgs) GetDealerId() int64 {
	return p.DealerId
}

func (p *AuctionServiceBiddingArgs) GetOrderId() int64 {
	return p.OrderId
}

func (p *AuctionServiceBiddingArgs) GetPrice() float64 {
	return p.Price
}
func (p *AuctionServiceBiddingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceBiddingArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DealerId = v
	}
	return nil
}

func (p *AuctionServiceBiddingArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.OrderId = v
	}
	return nil
}

func (p *AuctionServiceBiddingArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Price = v
	}
	return nil
}

func (p *AuctionServiceBiddingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bidding_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceBiddingArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("dealerId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dealerId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.DealerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.dealerId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dealerId: ", p), err)
	}
	return err
}

func (p *AuctionServiceBiddingArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("orderId", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:orderId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.OrderId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.orderId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:orderId: ", p), err)
	}
	return err
}

func (p *AuctionServiceBiddingArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("price", thrift.DOUBLE, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:price: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Price)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.price (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:price: ", p), err)
	}
	return err
}

func (p *AuctionServiceBiddingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceBiddingArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Ex
type AuctionServiceBiddingResult struct {
	Success *bool             `thrift:"success,0" json:"success,omitempty"`
	Ex      *InvalidException `thrift:"ex,1" json:"ex,omitempty"`
}

func NewAuctionServiceBiddingResult() *AuctionServiceBiddingResult {
	return &AuctionServiceBiddingResult{}
}

var AuctionServiceBiddingResult_Success_DEFAULT bool

func (p *AuctionServiceBiddingResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return AuctionServiceBiddingResult_Success_DEFAULT
	}
	return *p.Success
}

var AuctionServiceBiddingResult_Ex_DEFAULT *InvalidException

func (p *AuctionServiceBiddingResult) GetEx() *InvalidException {
	if !p.IsSetEx() {
		return AuctionServiceBiddingResult_Ex_DEFAULT
	}
	return p.Ex
}
func (p *AuctionServiceBiddingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuctionServiceBiddingResult) IsSetEx() bool {
	return p.Ex != nil
}

func (p *AuctionServiceBiddingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceBiddingResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AuctionServiceBiddingResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = &InvalidException{}
	if err := p.Ex.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ex), err)
	}
	return nil
}

func (p *AuctionServiceBiddingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bidding_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceBiddingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceBiddingResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetEx() {
		if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ex: ", p), err)
		}
		if err := p.Ex.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ex), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ex: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceBiddingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceBiddingResult(%+v)", *p)
}

// Attributes:
//  - DealerId
//  - OrderId
//  - Price
type AuctionServiceBidArgs struct {
	DealerId int64   `thrift:"dealerId,1" json:"dealerId"`
	OrderId  int64   `thrift:"orderId,2" json:"orderId"`
	Price    float64 `thrift:"price,3" json:"price"`
}

func NewAuctionServiceBidArgs() *AuctionServiceBidArgs {
	return &AuctionServiceBidArgs{}
}

func (p *AuctionServiceBidArgs) GetDealerId() int64 {
	return p.DealerId
}

func (p *AuctionServiceBidArgs) GetOrderId() int64 {
	return p.OrderId
}

func (p *AuctionServiceBidArgs) GetPrice() float64 {
	return p.Price
}
func (p *AuctionServiceBidArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceBidArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DealerId = v
	}
	return nil
}

func (p *AuctionServiceBidArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.OrderId = v
	}
	return nil
}

func (p *AuctionServiceBidArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Price = v
	}
	return nil
}

func (p *AuctionServiceBidArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bid_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceBidArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("dealerId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dealerId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.DealerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.dealerId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dealerId: ", p), err)
	}
	return err
}

func (p *AuctionServiceBidArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("orderId", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:orderId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.OrderId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.orderId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:orderId: ", p), err)
	}
	return err
}

func (p *AuctionServiceBidArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("price", thrift.DOUBLE, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:price: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Price)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.price (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:price: ", p), err)
	}
	return err
}

func (p *AuctionServiceBidArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceBidArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Ex
type AuctionServiceBidResult struct {
	Success *bool             `thrift:"success,0" json:"success,omitempty"`
	Ex      *InvalidException `thrift:"ex,1" json:"ex,omitempty"`
}

func NewAuctionServiceBidResult() *AuctionServiceBidResult {
	return &AuctionServiceBidResult{}
}

var AuctionServiceBidResult_Success_DEFAULT bool

func (p *AuctionServiceBidResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return AuctionServiceBidResult_Success_DEFAULT
	}
	return *p.Success
}

var AuctionServiceBidResult_Ex_DEFAULT *InvalidException

func (p *AuctionServiceBidResult) GetEx() *InvalidException {
	if !p.IsSetEx() {
		return AuctionServiceBidResult_Ex_DEFAULT
	}
	return p.Ex
}
func (p *AuctionServiceBidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuctionServiceBidResult) IsSetEx() bool {
	return p.Ex != nil
}

func (p *AuctionServiceBidResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceBidResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AuctionServiceBidResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = &InvalidException{}
	if err := p.Ex.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ex), err)
	}
	return nil
}

func (p *AuctionServiceBidResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bid_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceBidResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceBidResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetEx() {
		if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ex: ", p), err)
		}
		if err := p.Ex.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ex), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ex: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceBidResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceBidResult(%+v)", *p)
}

// Attributes:
//  - Scenne
type AuctionServiceStartAuctionArgs struct {
	Scenne *Scene `thrift:"scenne,1" json:"scenne"`
}

func NewAuctionServiceStartAuctionArgs() *AuctionServiceStartAuctionArgs {
	return &AuctionServiceStartAuctionArgs{}
}

var AuctionServiceStartAuctionArgs_Scenne_DEFAULT *Scene

func (p *AuctionServiceStartAuctionArgs) GetScenne() *Scene {
	if !p.IsSetScenne() {
		return AuctionServiceStartAuctionArgs_Scenne_DEFAULT
	}
	return p.Scenne
}
func (p *AuctionServiceStartAuctionArgs) IsSetScenne() bool {
	return p.Scenne != nil
}

func (p *AuctionServiceStartAuctionArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionArgs) readField1(iprot thrift.TProtocol) error {
	p.Scenne = &Scene{}
	if err := p.Scenne.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Scenne), err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("startAuction_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("scenne", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:scenne: ", p), err)
	}
	if err := p.Scenne.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Scenne), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:scenne: ", p), err)
	}
	return err
}

func (p *AuctionServiceStartAuctionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceStartAuctionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Ex
type AuctionServiceStartAuctionResult struct {
	Success *bool             `thrift:"success,0" json:"success,omitempty"`
	Ex      *InvalidException `thrift:"ex,1" json:"ex,omitempty"`
}

func NewAuctionServiceStartAuctionResult() *AuctionServiceStartAuctionResult {
	return &AuctionServiceStartAuctionResult{}
}

var AuctionServiceStartAuctionResult_Success_DEFAULT bool

func (p *AuctionServiceStartAuctionResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return AuctionServiceStartAuctionResult_Success_DEFAULT
	}
	return *p.Success
}

var AuctionServiceStartAuctionResult_Ex_DEFAULT *InvalidException

func (p *AuctionServiceStartAuctionResult) GetEx() *InvalidException {
	if !p.IsSetEx() {
		return AuctionServiceStartAuctionResult_Ex_DEFAULT
	}
	return p.Ex
}
func (p *AuctionServiceStartAuctionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuctionServiceStartAuctionResult) IsSetEx() bool {
	return p.Ex != nil
}

func (p *AuctionServiceStartAuctionResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AuctionServiceStartAuctionResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = &InvalidException{}
	if err := p.Ex.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ex), err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("startAuction_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuctionServiceStartAuctionResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceStartAuctionResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetEx() {
		if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ex: ", p), err)
		}
		if err := p.Ex.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ex), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ex: ", p), err)
		}
	}
	return err
}

func (p *AuctionServiceStartAuctionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuctionServiceStartAuctionResult(%+v)", *p)
}