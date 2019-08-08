package main

import (
	// Import the entire framework (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	// Use the new AXI protocol package
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

// Magic identifier for exporting
func Top(
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Disable AXI memory accesses.
	go axiprotocol.ReadDisable(memReadAddr, memReadData)
	go axiprotocol.WriteDisable(memWriteAddr, memWriteData, memWriteResp)
}
