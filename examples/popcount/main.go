package main

import (
	"teak"

	// Import support for AXI memory access
	aximemory "axi/memory"
	axiprotocol "axi/protocol"
	// Import the entire framework (including bundled verilog)
	_ "sdaccel"
)

// This is a kernel for showcasing the builtin Smash functionality.
// It implements popcount, by counting the number of 1s in a
// uint32. We're going to transform it into a bool array, and loop
// through counting each 1.

// Magic identifier for exporting
func Top(
	a uint32,
	output uintptr,

	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Disable memory reads
	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	// This will create a [32]bool
	var array [32]bool = teak.SmashU32(a)

	// We'll store results in sum
	sum := uint32(0)

	// We can now access each individual bit as a boolean
	for i := 0; i < 32; i++ {
		if array[i] {
			sum += 1
		}
	}

	// Write out the result.
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, output, sum)

}
