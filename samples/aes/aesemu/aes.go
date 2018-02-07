package main

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"

	"flag"

	"runtime/debug"

	"gitlab.com/yaotsu/core"
	"gitlab.com/yaotsu/core/connections"
	"gitlab.com/yaotsu/core/engines"
	"gitlab.com/yaotsu/core/util"
	"gitlab.com/yaotsu/gcn3"
	"gitlab.com/yaotsu/gcn3/emu"
	"gitlab.com/yaotsu/gcn3/insts"
	"gitlab.com/yaotsu/gcn3/kernels"
	"gitlab.com/yaotsu/mem"
)


type hostComponent struct {
   *core.ComponentBase
}

func newHostComponent() *hostComponent {
	h := new(hostComponent)
	h.ComponentBase = core.NewComponentBase( "host")
	h.AddPort("ToGpu")
	return h
}

func (h *hostComponent) Recv(req core.Req) *core.Error {
	switch req.(type) {
	case *kernels.LaunchKernelReq:
		log.Println("Kernel completed.")
	}
	return nil
}

func (h *hostComponent) Handle(evt core.Event) error {
	return nil
}

var (
	engine     core.Engine
	globalMem  *mem.IdealMemController
	gpu        *gcn3.Gpu
	host       *hostComponent
	connection core.Connection
	hsaco      *insts.HsaCo
	logger     *log.Logger
)


var cpuprofile = flag.String("cpuprofile", "prof.prof", "write cpu profile to file")
var kernel = flag.String("kernel", "../disasm/kernels.hsaco", "the kernel hsaco file")


func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	runtime.SetBlockProfileRate(1)
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		debug.PrintStack()
		os.Exit(1)
	}()

	// log.SetOutput(ioutil.Discard)
	logger = log.New(os.Stdout, "", 0)

	initPlatform()
	loadProgram()
	initMem()
	run()
	checkResult()
}


func initPlatform() {
	// Simulation engine
	engine = engines.NewSerialEngine()
	// engine.AcceptHook(core.NewLogEventHook(log.New(os.Stdout, "", 0)))

	// Connection
	connection = connections.NewDirectConnection(engine)

	// Memory
	globalMem = mem.NewIdealMemController("GlobalMem", engine, 4*mem.GB)
	globalMem.Freq = 1 * util.GHz
	globalMem.Latency = 1

	// Host
	host = newHostComponent()

	// Gpu
	gpu = gcn3.NewGpu("GPU")
	commandProcessor := gcn3.NewCommandProcessor("GPU.CommandProcessor")

	dispatcher := gcn3.NewDispatcher("GPU.Dispatcher", engine,
		new(kernels.GridBuilderImpl))
	dispatcher.Freq = 1 * util.GHz
	wgCompleteLogger := new(gcn3.WGCompleteLogger)
	wgCompleteLogger.Logger = logger
	dispatcher.AcceptHook(wgCompleteLogger)

	gpu.CommandProcessor = commandProcessor
	gpu.Driver = host
	commandProcessor.Dispatcher = dispatcher
	commandProcessor.Driver = gpu
	disassembler := insts.NewDisassembler()
	isaDebug, err := os.Create("isa.debug")
	if err != nil {
		fmt.Print("Isa debug file failed to open\n")
	}
	for i := 0; i < 4; i++ {
		scratchpadPreparer := emu.NewScratchpadPreparerImpl()
		alu := emu.NewALU(globalMem.Storage)
		computeUnit := emu.NewComputeUnit(fmt.Sprintf("%s.cu%d", gpu.Name(), i),
			engine, disassembler, scratchpadPreparer, alu)
		computeUnit.Freq = 1 * util.GHz
		computeUnit.GlobalMemStorage = globalMem.Storage
		dispatcher.CUs = append(dispatcher.CUs, computeUnit)
		core.PlugIn(computeUnit, "ToDispatcher", connection)

		wfHook := emu.NewWfHook(log.New(isaDebug, "", 0))
		computeUnit.AcceptHook(wfHook)
	}

	// Connection
	core.PlugIn(gpu, "ToCommandProcessor", connection)
	core.PlugIn(gpu, "ToDriver", connection)
	core.PlugIn(commandProcessor, "ToDriver", connection)
	core.PlugIn(commandProcessor, "ToDispatcher", connection)
	core.PlugIn(host, "ToGpu", connection)
	core.PlugIn(dispatcher, "ToCommandProcessor", connection)
	core.PlugIn(dispatcher, "ToCUs", connection)
	core.PlugIn(globalMem, "Top", connection)
}

func loadProgram() {
	executable, err := elf.Open(*kernel)
	if err != nil {
		log.Fatal(err)
	}

	sec := executable.Section(".text")
	hsacoData, err := sec.Data()
	if err != nil {
		log.Fatal(err)
	}

	err = globalMem.Storage.Write(0, hsacoData)
	if err != nil {
		log.Fatal(err)
	}

	hsaco = insts.NewHsaCoFromData(hsacoData)
	fmt.Println(hsaco.Info())
}

func initMem() {
	// Write the key
	key := []uint{2775827103,2708915352,2843725459,2775761052,374450381,38059738,442344641,104905438,
		2928140272,267459432,2794378747,66852199,1843523912,1873104786,1979247443,1941459341,3327558271,
		3383204119,1864977644,1825921419,1038236277,1380414951,666869428,1409797945,199004255,3262843208,
		2907850148,3246857263,1173726816,397595527,806178099,1678410250,2094003996,3199532628,333888496,
		3529615327,4028300030,3886557561,3617940554,3014649408,625081969,2616524837,2282994645,1517427722,
		1314547353,2851229664,2119642026,3455634922,620526028,3205069289,924500540,1835589174}
	keybuffer := make([]byte,0)
	buffer := bytes.NewBuffer(keybuffer)
	binary.Write(buffer, binary.LittleEndian,key)

	err := globalMem.Storage.Write(4*mem.KB, buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// Write the input
	inputData := make([]byte, 0)
	buffer = bytes.NewBuffer(inputData)
	for i := 0; i < 32768; i++ {
		binary.Write(buffer, binary.LittleEndian, uint8(i))
	}
	err = globalMem.Storage.Write(8*mem.KB, buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}

func run() {
	kernelArgsBuffer := bytes.NewBuffer(make([]byte, 0))

	binary.Write(kernelArgsBuffer, binary.LittleEndian, uint64(8192)) // Input
	binary.Write(kernelArgsBuffer, binary.LittleEndian, uint64(4096)) // Key

	err := globalMem.Storage.Write(65536, kernelArgsBuffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	req := kernels.NewLaunchKernelReq()
	req.HsaCo = hsaco
	req.Packet = new(kernels.HsaKernelDispatchPacket)
	req.Packet.GridSizeX = 256 * 4
	req.Packet.GridSizeY = 1
	req.Packet.GridSizeZ = 1
	req.Packet.WorkgroupSizeX = 256
	req.Packet.WorkgroupSizeY = 1
	req.Packet.WorkgroupSizeZ = 1
	req.Packet.KernelObject = 0
	req.Packet.KernargAddress = 65536

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, req.Packet)
	err = globalMem.Storage.Write(0x11000, buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	req.PacketAddress = 0x11000
	req.SetSrc(host)
	req.SetDst(gpu)
	req.SetSendTime(0)
	connErr := connection.Send(req)
	if connErr != nil {
		log.Fatal(connErr)
	}

	engine.Run()
}

func checkResult() {
	buf, err := globalMem.Storage.Read(12*mem.KB, 1024*4)
	if err != nil {
		log.Fatal(nil)
	}

	for i := 0; i < 1024; i++ {
		bits := binary.LittleEndian.Uint32(buf[i*4 : i*4+4])
		filtered := math.Float32frombits(bits)

		fmt.Printf("%d: %f\n", i, filtered)
	}
}



