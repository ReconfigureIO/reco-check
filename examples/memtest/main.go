package main

import (
	// import the entire framework (including bundled verilog)
	_ "sdaccel"
	"sdaccel/memory"
	// Use the new AXI protocol package
	aximemory "axi/memory"
	axiprotocol "axi/protocol"
)

// magic identifier for exporting
func Top(
	outputData uintptr,

	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	go memory.DisableReads(memReadAddr, memReadData)

	for out := uint64(0); out < 4294967295; out += 4 {
		aximemory.WriteUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputData, 1)
	}

}
