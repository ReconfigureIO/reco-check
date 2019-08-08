//
// (c) 2017 ReconfigureIO
//
// <COPYRIGHT TERMS>
//

//
// AXI legacy interface to memory mapped RAM and I/O. This has been superseded
// by the new AXI protocol and memory access packages, but is retained for
// compatibility with existing demonstration and test code.
// TODO: Remove this once all code has been ported to the new API.
//

package memory

import (
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

// Goroutine to disable memory bus read transactions. Should only be run
// once for each memory interface.
func DisableReads(memoryReadAddr chan<- axiprotocol.Addr,
	memoryReadData <-chan axiprotocol.ReadData) {
	axiprotocol.ReadDisable(memoryReadAddr, memoryReadData)
}

// Goroutine to disable memory bus write transactions. Should only be run once
// for each memory interface.
func DisableWrites(
	memoryWriteAddr chan<- axiprotocol.Addr,
	memoryWriteData chan<- axiprotocol.WriteData,
	memoryWriteResp <-chan axiprotocol.WriteResp) {
	axiprotocol.WriteDisable(memoryWriteAddr, memoryWriteData, memoryWriteResp)
}

// NOTE: API change replaces returned response with boolean 'writeOk' flag.
func Write(
	address uintptr,
	data uint32,
	memoryWriteAddr chan<- axiprotocol.Addr,
	memoryWriteData chan<- axiprotocol.WriteData,
	memoryWriteResp <-chan axiprotocol.WriteResp) bool {
	return aximemory.WriteUInt32(memoryWriteAddr, memoryWriteData,
		memoryWriteResp, true, address, data)
}

func Read(
	address uintptr,
	memoryReadAddr chan<- axiprotocol.Addr,
	memoryReadData <-chan axiprotocol.ReadData) uint32 {
	return aximemory.ReadUInt32(memoryReadAddr, memoryReadData,
		true, address)
}

func ReadBurst(
	address uintptr,
	burst uint32,
	dataStream chan<- uint32,
	memoryReadAddr chan<- axiprotocol.Addr,
	memoryReadData <-chan axiprotocol.ReadData) {
	aximemory.ReadBurstUInt32(memoryReadAddr, memoryReadData,
		true, address, burst, dataStream)
}

func WriteBurst(
	address uintptr,
	burst uint32,
	dataStream <-chan uint32,
	memoryWriteAddr chan<- axiprotocol.Addr,
	memoryWriteData chan<- axiprotocol.WriteData,
	memoryWriteResp <-chan axiprotocol.WriteResp) {
	aximemory.WriteBurstUInt32(memoryWriteAddr, memoryWriteData,
		memoryWriteResp, true, address, burst, dataStream)
}
