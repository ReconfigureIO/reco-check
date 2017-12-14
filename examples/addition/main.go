package main

import (
	// Import the entire framework (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	// Use the new AXI protocol package
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

// Magic identifier for exporting
func Top(
	a uint32,
	b uint32,
	addr uintptr,

	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Disable memory reads
	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	val := a + b

	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
