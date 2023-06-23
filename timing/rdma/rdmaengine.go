//Package rdma provides the implementation of an RDMA engine.
package rdma

import (
	"log"
	"reflect"
	"fmt"
	"os"
	"gitlab.com/akita/akita/v3/sim"
	"gitlab.com/akita/akita/v3/tracing"
	"gitlab.com/akita/mem/v3/mem"
)

type transaction struct {
	fromInside  sim.Msg
	fromOutside sim.Msg
	toInside    sim.Msg
	toOutside   sim.Msg
	sendTime	sim.VTimeInSec
	recvTime	sim.VTimeInSec
	addr	 uint64
}

// An Engine is a component that helps one GPU to access the memory on
// another GPU
type Engine struct {
	*sim.TickingComponent

	ToOutside sim.Port

	ToL1 sim.Port
	ToL2 sim.Port

	CtrlPort sim.Port

	isDraining              bool
	pauseIncomingReqsFromL1 bool
	currentDrainReq         *DrainReq

	localModules           mem.LowModuleFinder
	RemoteRDMAAddressTable mem.LowModuleFinder

	transactionsFromOutside []transaction
	transactionsFromInside  []transaction
	transactionEpoch        sim.VTimeInSec
	transactionEpochRsp	sim.VTimeInSec
	srcUsageCount       map[string]int
	dstUsageCount       map[string]int
	dstGPU                  []string
}

// SetLocalModuleFinder sets the table to lookup for local data.
func (e *Engine) SetLocalModuleFinder(lmf mem.LowModuleFinder) {
	e.localModules = lmf
}

// Tick checks if make progress
func (e *Engine) Tick(now sim.VTimeInSec) bool {
	madeProgress := false

	madeProgress = e.processFromCtrlPort(now) || madeProgress
	if e.isDraining {
		madeProgress = e.drainRDMA(now) || madeProgress
	}
	madeProgress = e.processFromL1(now) || madeProgress
	madeProgress = e.processFromL2(now) || madeProgress
	madeProgress = e.processFromOutside(now) || madeProgress

	return madeProgress
}

func (e *Engine) processFromCtrlPort(now sim.VTimeInSec) bool {
	req := e.CtrlPort.Peek()
	if req == nil {
		return false
	}

	req = e.CtrlPort.Retrieve(now)
	switch req := req.(type) {
	case *DrainReq:
		e.currentDrainReq = req
		e.isDraining = true
		e.pauseIncomingReqsFromL1 = true
		return true
	case *RestartReq:
		return e.processRDMARestartReq(now)
	default:
		log.Panicf("cannot process request of type %s", reflect.TypeOf(req))
		return false
	}
}

func (e *Engine) processRDMARestartReq(now sim.VTimeInSec) bool {
	restartCompleteRsp := RestartRspBuilder{}.
		WithSendTime(now).
		WithSrc(e.CtrlPort).
		WithDst(e.currentDrainReq.Src).
		Build()
	err := e.CtrlPort.Send(restartCompleteRsp)

	if err != nil {
		return false
	}
	e.currentDrainReq = nil
	e.pauseIncomingReqsFromL1 = false

	return true
}

func (e *Engine) drainRDMA(now sim.VTimeInSec) bool {
	if e.fullyDrained() {
		drainCompleteRsp := DrainRspBuilder{}.
			WithSendTime(now).
			WithSrc(e.CtrlPort).
			WithDst(e.currentDrainReq.Src).
			Build()

		err := e.CtrlPort.Send(drainCompleteRsp)
		if err != nil {
			return false
		}
		e.isDraining = false
		return true
	}
	return false
}
//EUN
func (e *Engine) writeFileRdmaFromInside(req mem.AccessRsp, epoch sim.VTimeInSec) {
	f, err := os.OpenFile("rdma.log", os.O_APPEND|os.O_RDWR, 0755)
        if err != nil {
        // Handle the error, such as creating the file if it doesn't exist
       	if os.IsNotExist(err) {
        	f, err = os.Create("rdma.log")
                        if err != nil {
                        	log.Fatal(err)
                        }
             	} else {
                	log.Fatal(err)
                	}
        }
	defer f.Close()
	fmt.Fprintf(f, "---------------------------------------[Transactions Resolved]-------------------------------------------\n")
	fmt.Fprintf(f,"send time=%f\tsrc=\t%s\tdst=\t%s\tAddr= %s\tepoch time= %f\n\n",req.Meta().SendTime,req.Meta().Src.Name(),req.Meta().Dst.Name(),req.Meta().ID,epoch)
}


//ES
func (e *Engine) countRequestUsage(trans transaction) {
	src := trans.fromInside.Meta().Src.Name()
	dst := trans.toOutside.Meta().Dst.Name()

	// Count usage as src
	if _, ok := e.srcUsageCount[src]; ok {
		e.srcUsageCount[src]++
	} else {
		e.srcUsageCount[src] = 1
	}

	// Count usage as dst
	if _, ok := e.dstUsageCount[dst]; ok {
		e.dstUsageCount[dst]++
	} else {
		e.dstUsageCount[dst] = 1
	}
	 f, err := os.OpenFile("./rdma.log", os.O_APPEND|os.O_RDWR, 0755)
                if err != nil {
        // Handle the error, such as creating the file if it doesn't exist
                        if os.IsNotExist(err) {
                                f, err = os.Create("rdma.log")
                                if err != nil {
                                        log.Fatal(err)
                                }
                        } else {
                                log.Fatal(err)
                        }
                }
                defer f.Close()
                fmt.Fprintf(f,"[Request Info]\n")
                fmt.Fprintf(f,"Send time= %f\tsrc=\t%s\tdst=\t%s\tAddr= %d\n",trans.sendTime,trans.fromInside.Meta().Src.Name(),trans.toOutside.Meta().Dst.Name(),trans.addr)
		fmt.Fprintf(f,"\n[SRC usage Count]\n")
		for src, count := range e.srcUsageCount {
			fmt.Fprintf(f,"[%s]: %d\n", src, count)
		}
		fmt.Fprintf(f,"\n[DST usage Count]\n")
		for dst, count := range e.dstUsageCount {
			fmt.Fprintf(f,"[%s]: %d\n", dst, count)
		}

}

func (e *Engine) fullyDrained() bool {
	return len(e.transactionsFromOutside) == 0 &&
		len(e.transactionsFromInside) == 0
}

func (e *Engine) processFromL1(now sim.VTimeInSec) bool {
	if e.pauseIncomingReqsFromL1 {
		return false
	}

	req := e.ToL1.Peek()
	if req == nil {
		return false
	}

	switch req := req.(type) {
	case mem.AccessReq:
		return e.processReqFromL1(now, req)
	default:
		log.Panicf("cannot process request of type %s", reflect.TypeOf(req))
		return false
	}
}

func (e *Engine) processFromL2(now sim.VTimeInSec) bool {
	req := e.ToL2.Peek()
	if req == nil {
		return false
	}

	switch req := req.(type) {
	case mem.AccessRsp:
		return e.processRspFromL2(now, req)
	default:
		panic("unknown req type")
	}
}

func (e *Engine) processFromOutside(now sim.VTimeInSec) bool {
	req := e.ToOutside.Peek()
	if req == nil {
		return false
	}

	switch req := req.(type) {
	case mem.AccessReq:
		return e.processReqFromOutside(now, req)
	case mem.AccessRsp:
		return e.processRspFromOutside(now, req)
	default:
		log.Panicf("cannot process request of type %s", reflect.TypeOf(req))
		return false
	}
}

func (e *Engine) processReqFromL1(
	now sim.VTimeInSec,
	req mem.AccessReq,
) bool {
	dst := e.RemoteRDMAAddressTable.Find(req.GetAddress())

	if dst == e.ToOutside {
		panic("RDMA loop back detected")
	}

	cloned := e.cloneReq(req)
	cloned.Meta().Src = e.ToOutside
	cloned.Meta().Dst = dst
	cloned.Meta().SendTime = now
	e.transactionEpoch = now
	err := e.ToOutside.Send(cloned)
	if err == nil {
		e.ToL1.Retrieve(now)

		tracing.TraceReqReceive(req, e)
		tracing.TraceReqInitiate(cloned, e, tracing.MsgIDAtReceiver(req, e))

		trans := transaction{
			fromInside: req,
			toOutside:  cloned,
			sendTime: now,
			addr: req.GetAddress(),
		}
		e.transactionsFromInside = append(e.transactionsFromInside, trans)
		e.countRequestUsage(trans)

		return true
	}

	return false
}

func (e *Engine) processReqFromOutside(
	now sim.VTimeInSec,
	req mem.AccessReq,
) bool {
	dst := e.localModules.Find(req.GetAddress())

	cloned := e.cloneReq(req)
	cloned.Meta().Src = e.ToL2
	cloned.Meta().Dst = dst
	cloned.Meta().SendTime = now

	e.transactionEpochRsp = now
	err := e.ToL2.Send(cloned)
	if err == nil {
		e.ToOutside.Retrieve(now)

		tracing.TraceReqReceive(req, e)
		tracing.TraceReqInitiate(cloned, e, tracing.MsgIDAtReceiver(req, e))

		trans := transaction{
			fromOutside: req,
			toInside:    cloned,
		}
		e.transactionsFromOutside =
			append(e.transactionsFromOutside, trans)
		return true
	}
	return false
}

func (e *Engine) processRspFromL2(
	now sim.VTimeInSec,
	rsp mem.AccessRsp,
) bool {
	transactionIndex := e.findTransactionByRspToID(
		rsp.GetRspTo(), e.transactionsFromOutside)
	trans := e.transactionsFromOutside[transactionIndex]

	rspToOutside := e.cloneRsp(rsp, trans.fromOutside.Meta().ID)
	rspToOutside.Meta().SendTime = now
	rspToOutside.Meta().Src = e.ToOutside
	rspToOutside.Meta().Dst = trans.fromOutside.Meta().Src

	fmt.Printf("src= %s\t->dst= %s\n",rspToOutside.Meta().Src.Name(),rspToOutside.Meta().Dst.Name())
	e.transactionEpoch = 0	
	err := e.ToOutside.Send(rspToOutside)
	if err == nil {
		e.ToL2.Retrieve(now)

		tracing.TraceReqFinalize(trans.toInside, e)
		tracing.TraceReqComplete(trans.fromOutside, e)

		e.transactionsFromOutside =
			append(e.transactionsFromOutside[:transactionIndex],
				e.transactionsFromOutside[transactionIndex+1:]...)
		return true
	}
	return false
}

func (e *Engine) processRspFromOutside(
	now sim.VTimeInSec,
	rsp mem.AccessRsp,
) bool {
	transactionIndex := e.findTransactionByRspToID(
		rsp.GetRspTo(), e.transactionsFromInside)
	trans := e.transactionsFromInside[transactionIndex]
	trans.recvTime = now
	
	rspToInside := e.cloneRsp(rsp, trans.fromInside.Meta().ID)
	rspToInside.Meta().SendTime = now
	rspToInside.Meta().Src = e.ToL1
	rspToInside.Meta().Dst = trans.fromInside.Meta().Src
	
	e.transactionEpoch = now - e.transactionEpoch
	e.transactionEpochRsp = now - e.transactionEpochRsp
	fmt.Printf("transaction epoch between 2 and 4= %f\n",e.transactionEpochRsp)
	e.transactionEpochRsp = 0
	e.writeFileRdmaFromInside(rspToInside,e.transactionEpoch)
	e.transactionEpoch = 0
	err := e.ToL1.Send(rspToInside)
	if err == nil {
		e.ToOutside.Retrieve(now)

		tracing.TraceReqFinalize(trans.toOutside, e)
		tracing.TraceReqComplete(trans.fromInside, e)

		e.transactionsFromInside =
			append(e.transactionsFromInside[:transactionIndex],
				e.transactionsFromInside[transactionIndex+1:]...)
		return true
	}

	return false
}
func (e *Engine) findTransactionByRspToID(
	rspTo string,
	transactions []transaction,
) int {
	for i, trans := range transactions {
		if trans.toOutside != nil && trans.toOutside.Meta().ID == rspTo {
			return i
		}

		if trans.toInside != nil && trans.toInside.Meta().ID == rspTo {
			return i
		}
	}

	log.Panicf("transaction %s not found", rspTo)
	return 0
}

func (e *Engine) cloneReq(origin mem.AccessReq) mem.AccessReq {
	switch origin := origin.(type) {
	case *mem.ReadReq:
		read := mem.ReadReqBuilder{}.
			WithSendTime(origin.SendTime).
			WithSrc(origin.Src).
			WithDst(origin.Dst).
			WithAddress(origin.Address).
			WithByteSize(origin.AccessByteSize).
			Build()
		return read
	case *mem.WriteReq:
		write := mem.WriteReqBuilder{}.
			WithSendTime(origin.SendTime).
			WithSrc(origin.Src).
			WithDst(origin.Dst).
			WithAddress(origin.Address).
			WithData(origin.Data).
			WithDirtyMask(origin.DirtyMask).
			Build()
		return write
	default:
		log.Panicf("cannot clone request of type %s",
			reflect.TypeOf(origin))
	}
	return nil
}

func (e *Engine) cloneRsp(origin mem.AccessRsp, rspTo string) mem.AccessRsp {
	switch origin := origin.(type) {
	case *mem.DataReadyRsp:
		rsp := mem.DataReadyRspBuilder{}.
			WithSendTime(origin.SendTime).
			WithSrc(origin.Src).
			WithDst(origin.Dst).
			WithRspTo(rspTo).
			WithData(origin.Data).
			Build()
		return rsp
	case *mem.WriteDoneRsp:
		rsp := mem.WriteDoneRspBuilder{}.
			WithSendTime(origin.SendTime).
			WithSrc(origin.Src).
			WithDst(origin.Dst).
			WithRspTo(rspTo).
			Build()
		return rsp
	default:
		log.Panicf("cannot clone request of type %s",
			reflect.TypeOf(origin))
	}
	return nil
}

// SetFreq sets freq
func (e *Engine) SetFreq(freq sim.Freq) {
	e.TickingComponent.Freq = freq
}

// NewEngine creates new engine
func NewEngine(
	name string,
	engine sim.Engine,
	localModules mem.LowModuleFinder,
	remoteModules mem.LowModuleFinder,
) *Engine {
	e := new(Engine)
	e.TickingComponent = sim.NewTickingComponent(name, engine, 1*sim.GHz, e)
	e.localModules = localModules
	e.RemoteRDMAAddressTable = remoteModules

	e.srcUsageCount = make(map[string]int)
	e.dstUsageCount = make(map[string]int)
	e.ToL1 = sim.NewLimitNumMsgPort(e, 1, name+".ToL1")
	e.ToL2 = sim.NewLimitNumMsgPort(e, 1, name+".ToL2")
	e.CtrlPort = sim.NewLimitNumMsgPort(e, 1, name+".CtrlPort")
	e.ToOutside = sim.NewLimitNumMsgPort(e, 1, name+".ToOutside")

	return e
}
