package runner

import (
	"fmt"
	"gitlab.com/akita/akita/v3/tracing"
	"gitlab.com/akita/mem/v3/vm"
	"strconv"
)

// tlbTracer can trace TLB access information.
type tlbTracer struct {
	traceWriter *tracing.CSVTraceWriter
}

var taskCount uint64 = 0

// newTLBTracer creates a tracer for TLB access.
func newTlbTracer(traceWriter *tracing.CSVTraceWriter) *tlbTracer {
	return &tlbTracer{
		traceWriter: traceWriter,
	}
}

func (t *tlbTracer) StartTask(task tracing.Task) {
	// Do nothing
}

func (t *tlbTracer) StepTask(task tracing.Task) {
	// Do nothing
}

func (t *tlbTracer) EndTask(task tracing.Task) {
	// Do nothing
}

// TraceTLBAccess records the TLB access information.
func (t *tlbTracer) TraceTLBAccess(
	req *vm.TranslationReq,
	setID int,
	wayID int,
	vAddr uint64,
	traceWriter *tracing.CSVTraceWriter,
) {
	taskCount++
	taskID := fmt.Sprintf("TLB-Access-%d", taskCount)

	task := tracing.Task{
		ID:        taskID,
		ParentID:  "",
		Kind:      "TLB Access",
		What:      "TLB Access",
		Where:     "",
		StartTime: 0,
		EndTime:   0,
	}
	traceWriter.Write(task)

	lookupTaskID := fmt.Sprintf("%s-Lookup", taskID)
	traceWriter.Write(tracing.Task{
		ID:        lookupTaskID,
		ParentID:  taskID,
		Kind:      "Lookup",
		What:      strconv.FormatUint(req.VAddr, 10),
		Where:     fmt.Sprintf("SetID: %d, WayID: %d", setID, wayID),
		StartTime: 0,
		EndTime:   0,
	})
}
