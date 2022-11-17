package main

import (
	"golangchallenge/processors"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessor_StartProcessing(t *testing.T) {
	wg := &sync.WaitGroup{}
	td := GetTripsData(wg)
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	p := processors.CreateProcessorFromData(td, wg)
	p.StartProcessing()

	// Check that the processor has finished processing
	assert.NotNil(t, p.GetTopRankedDriver())
	assert.NotNil(t, p.GetTopRankedHotel())
	runtime.ReadMemStats(&m2)

	assert.LessOrEqual(t, m2.Alloc-m1.Alloc, uint64(128*1024*1024))
}
