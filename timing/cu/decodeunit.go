package cu

import (
	"log"
	"os"
	"fmt"

	"gitlab.com/akita/akita/v3/sim"
	"gitlab.com/akita/mgpusim/v3/timing/wavefront"
)

// A DecodeUnit is any type of decode unit that takes one cycle to decode
type DecodeUnit struct {
	cu        *ComputeUnit
	ExecUnits []SubComponent // Execution units, index by SIMD number

	toDecode *wavefront.Wavefront
	decoded  bool

	isIdle bool
}

// NewDecodeUnit creates a new decode unit
func NewDecodeUnit(cu *ComputeUnit) *DecodeUnit {
	du := new(DecodeUnit)
	du.cu = cu
	du.decoded = false
	return du
}

// AddExecutionUnit registers an executions unit to the decode unit, so that
// the decode unit knows where to send the instruction to after decoding.
// This function has to be called in the order of SIMD number.
func (du *DecodeUnit) AddExecutionUnit(cuComponent SubComponent) {
	du.ExecUnits = append(du.ExecUnits, cuComponent)
}

// CanAcceptWave checks if the DecodeUnit is ready to decode another
// instruction
func (du *DecodeUnit) CanAcceptWave() bool {
	return du.toDecode == nil
}

// IsIdle checks idleness
func (du *DecodeUnit) IsIdle() bool {
	du.isIdle = (du.toDecode == nil) && (du.decoded == false)
	return du.isIdle
}

// AcceptWave takes a wavefront and decode the instruction in the next cycle
func (du *DecodeUnit) AcceptWave(
	wave *wavefront.Wavefront,
	now sim.VTimeInSec,
) {
	if du.toDecode != nil {
		log.Panicf("Decode unit busy, please run CanAcceptWave before accepting a wave")
	}

	du.toDecode = wave
	f, err := os.OpenFile("./instDecoding.log", os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		// Handle the error, such as creating the file if it doesn't exist
		if os.IsNotExist(err) {
			f, err = os.Create("instDecoding.log")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	defer f.Close()
	fmt.Fprintf(f, "[Inst Name]=\t%s\t[Inst ID]=\t%d\t[Format Name]=\t%s\n",du.toDecode.Inst().InstName,du.toDecode.Inst().ID,du.toDecode.Inst().FormatName)
	du.decoded = false
}

// Run decodes the instruction and sends the instruction to the next pipeline
// stage
func (du *DecodeUnit) Run(now sim.VTimeInSec) bool {
	if du.toDecode != nil {
		simdID := du.toDecode.SIMDID
		execUnit := du.ExecUnits[simdID]

		if execUnit.CanAcceptWave() {
			execUnit.AcceptWave(du.toDecode, now)
			du.toDecode = nil
			return true
		}
	}

	if du.toDecode != nil && !du.decoded {
		du.decoded = true
		return true
	}

	return false
}

// Flush clear the unit
func (du *DecodeUnit) Flush() {
	du.toDecode = nil
}
