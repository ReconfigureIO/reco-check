//
// (c) 2017 ReconfigureIO
//
// <COPYRIGHT TERMS>
//

//
// AXI protocol bus arbitration between multiple 'upstream' ports. This package
// specifies a set of goroutines which may be used to arbitrate between multiple
// upstream AXI 'server' ports and a single downstream 'client' port. The
// current implementation supports arbitration between 2, 3 or 4 upstream ports.
// TODO: Support arbitrary number of upstream ports on demand using the Go
// generate capability.
//

/*
Package arbitrate provides reusable arbitrators for AXI transations.
*/
package arbitrate

import (
	"axi/protocol"
)

//
// Perform write data transfers.
//
func writeDataTransfer(
	clientData chan<- protocol.WriteData,
	serverData <-chan protocol.WriteData) {

	// Terminate transfers on write data channel 'last' flag.
	getNext := true
	for getNext {
		writeData := <-serverData
		clientData <- writeData
		getNext = !writeData.Last
	}
}

//
// Perform read data transfers.
//
func readDataTransfer(
	clientData <-chan protocol.ReadData,
	serverData chan<- protocol.ReadData) {

	// Terminate transfers on write data channel 'last' flag.
	getNext := true
	for getNext {
		readData := <-clientData
		serverData <- readData
		getNext = !readData.Last
	}
}

//
// Goroutine which implements AXI arbitration between two AXI write interfaces.
//
func WriteArbitrateX2(
	clientAddr chan<- protocol.Addr,
	clientData chan<- protocol.WriteData,
	clientResp <-chan protocol.WriteResp,
	serverAddr0 <-chan protocol.Addr,
	serverData0 <-chan protocol.WriteData,
	serverResp0 chan<- protocol.WriteResp,
	serverAddr1 <-chan protocol.Addr,
	serverData1 <-chan protocol.WriteData,
	serverResp1 chan<- protocol.WriteResp) {

	// Specify the input selection channels.
	dataChanSelect := make(chan byte)
	respChanSelect := make(chan byte)

	// Run write data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			respChanSelect <- chanSelect
			switch chanSelect {
			case 0:
				writeDataTransfer(clientData, serverData0)
			default:
				writeDataTransfer(clientData, serverData1)
			}
		}
	}()

	// Run response channel handler.
	go func() {
		for {
			chanSelect := <-respChanSelect
			writeResp := <-clientResp
			switch chanSelect {
			case 0:
				serverResp0 <- writeResp
			default:
				serverResp1 <- writeResp
			}
		}
	}()

	// Use independent write addresses for efficient implementation.
	var writeAddr0 protocol.Addr
	var writeAddr1 protocol.Addr
	for {
		select {
		case writeAddr0 = <-serverAddr0:
			clientAddr <- writeAddr0
			dataChanSelect <- 0
		case writeAddr1 = <-serverAddr1:
			clientAddr <- writeAddr1
			dataChanSelect <- 1
		}
	}
}

//
// Goroutine which implements AXI arbitration between three AXI write interfaces.
//
func WriteArbitrateX3(
	clientAddr chan<- protocol.Addr,
	clientData chan<- protocol.WriteData,
	clientResp <-chan protocol.WriteResp,
	serverAddr0 <-chan protocol.Addr,
	serverData0 <-chan protocol.WriteData,
	serverResp0 chan<- protocol.WriteResp,
	serverAddr1 <-chan protocol.Addr,
	serverData1 <-chan protocol.WriteData,
	serverResp1 chan<- protocol.WriteResp,
	serverAddr2 <-chan protocol.Addr,
	serverData2 <-chan protocol.WriteData,
	serverResp2 chan<- protocol.WriteResp) {

	// Specify the input selection channels.
	dataChanSelect := make(chan byte)
	respChanSelect := make(chan byte)

	// Run write data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			respChanSelect <- chanSelect
			switch chanSelect {
			case 0:
				writeDataTransfer(clientData, serverData0)
			case 1:
				writeDataTransfer(clientData, serverData1)
			default:
				writeDataTransfer(clientData, serverData2)
			}
		}
	}()

	// Run response channel handler.
	go func() {
		for {
			chanSelect := <-respChanSelect
			writeResp := <-clientResp
			switch chanSelect {
			case 0:
				serverResp0 <- writeResp
			case 1:
				serverResp1 <- writeResp
			default:
				serverResp2 <- writeResp
			}
		}
	}()

	// Use independent write addresses for efficient implementation.
	var writeAddr0 protocol.Addr
	var writeAddr1 protocol.Addr
	var writeAddr2 protocol.Addr
	for {
		select {
		case writeAddr0 = <-serverAddr0:
			clientAddr <- writeAddr0
			dataChanSelect <- 0
		case writeAddr1 = <-serverAddr1:
			clientAddr <- writeAddr1
			dataChanSelect <- 1
		case writeAddr2 = <-serverAddr2:
			clientAddr <- writeAddr2
			dataChanSelect <- 2
		}
	}
}

//
// Goroutine which implements AXI arbitration between four AXI write interfaces.
//
func WriteArbitrateX4(
	clientAddr chan<- protocol.Addr,
	clientData chan<- protocol.WriteData,
	clientResp <-chan protocol.WriteResp,
	serverAddr0 <-chan protocol.Addr,
	serverData0 <-chan protocol.WriteData,
	serverResp0 chan<- protocol.WriteResp,
	serverAddr1 <-chan protocol.Addr,
	serverData1 <-chan protocol.WriteData,
	serverResp1 chan<- protocol.WriteResp,
	serverAddr2 <-chan protocol.Addr,
	serverData2 <-chan protocol.WriteData,
	serverResp2 chan<- protocol.WriteResp,
	serverAddr3 <-chan protocol.Addr,
	serverData3 <-chan protocol.WriteData,
	serverResp3 chan<- protocol.WriteResp) {

	// Specify the input selection channels.
	dataChanSelect := make(chan byte)
	respChanSelect := make(chan byte)

	// Run write data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			respChanSelect <- chanSelect
			switch chanSelect {
			case 0:
				writeDataTransfer(clientData, serverData0)
			case 1:
				writeDataTransfer(clientData, serverData1)
			case 2:
				writeDataTransfer(clientData, serverData2)
			default:
				writeDataTransfer(clientData, serverData3)
			}
		}
	}()

	// Run response channel handler.
	go func() {
		for {
			chanSelect := <-respChanSelect
			writeResp := <-clientResp
			switch chanSelect {
			case 0:
				serverResp0 <- writeResp
			case 1:
				serverResp1 <- writeResp
			case 2:
				serverResp2 <- writeResp
			default:
				serverResp3 <- writeResp
			}
		}
	}()

	// Use independent write addresses for efficient implementation.
	var writeAddr0 protocol.Addr
	var writeAddr1 protocol.Addr
	var writeAddr2 protocol.Addr
	var writeAddr3 protocol.Addr
	for {
		select {
		case writeAddr0 = <-serverAddr0:
			clientAddr <- writeAddr0
			dataChanSelect <- 0
		case writeAddr1 = <-serverAddr1:
			clientAddr <- writeAddr1
			dataChanSelect <- 1
		case writeAddr2 = <-serverAddr2:
			clientAddr <- writeAddr2
			dataChanSelect <- 2
		case writeAddr3 = <-serverAddr3:
			clientAddr <- writeAddr3
			dataChanSelect <- 3
		}
	}
}

//
// Goroutine which implements AXI arbitration between two AXI read interfaces.
//
func ReadArbitrateX2(
	clientAddr chan<- protocol.Addr,
	clientData <-chan protocol.ReadData,
	serverAddr0 <-chan protocol.Addr,
	serverData0 chan<- protocol.ReadData,
	serverAddr1 <-chan protocol.Addr,
	serverData1 chan<- protocol.ReadData) {

	// Specify the input selection channel.
	dataChanSelect := make(chan byte)

	// Run read data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			switch chanSelect {
			case 0:
				readDataTransfer(clientData, serverData0)
			default:
				readDataTransfer(clientData, serverData1)
			}
		}
	}()

	// Use independent read addresses for efficient implementation.
	var readAddr0 protocol.Addr
	var readAddr1 protocol.Addr
	for {
		select {
		case readAddr0 = <-serverAddr0:
			clientAddr <- readAddr0
			dataChanSelect <- 0
		case readAddr1 = <-serverAddr1:
			clientAddr <- readAddr1
			dataChanSelect <- 1
		}
	}
}

//
// Goroutine which implements AXI arbitration between three AXI read interfaces.
//
func ReadArbitrateX3(
	clientAddr chan<- protocol.Addr,
	clientData <-chan protocol.ReadData,
	serverAddr0 <-chan protocol.Addr,
	serverData0 chan<- protocol.ReadData,
	serverAddr1 <-chan protocol.Addr,
	serverData1 chan<- protocol.ReadData,
	serverAddr2 <-chan protocol.Addr,
	serverData2 chan<- protocol.ReadData) {

	// Specify the input selection channel.
	dataChanSelect := make(chan byte)

	// Run read data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			switch chanSelect {
			case 0:
				readDataTransfer(clientData, serverData0)
			case 1:
				readDataTransfer(clientData, serverData1)
			default:
				readDataTransfer(clientData, serverData2)
			}
		}
	}()

	// Use independent read addresses for efficient implementation.
	var readAddr0 protocol.Addr
	var readAddr1 protocol.Addr
	var readAddr2 protocol.Addr
	for {
		select {
		case readAddr0 = <-serverAddr0:
			clientAddr <- readAddr0
			dataChanSelect <- 0
		case readAddr1 = <-serverAddr1:
			clientAddr <- readAddr1
			dataChanSelect <- 1
		case readAddr2 = <-serverAddr2:
			clientAddr <- readAddr2
			dataChanSelect <- 2
		}
	}
}

//
// Goroutine which implements AXI arbitration between four AXI read interfaces.
//
func ReadArbitrateX4(
	clientAddr chan<- protocol.Addr,
	clientData <-chan protocol.ReadData,
	serverAddr0 <-chan protocol.Addr,
	serverData0 chan<- protocol.ReadData,
	serverAddr1 <-chan protocol.Addr,
	serverData1 chan<- protocol.ReadData,
	serverAddr2 <-chan protocol.Addr,
	serverData2 chan<- protocol.ReadData,
	serverAddr3 <-chan protocol.Addr,
	serverData3 chan<- protocol.ReadData) {

	// Specify the input selection channel.
	dataChanSelect := make(chan byte)

	// Run read data channel handler.
	go func() {
		for {
			chanSelect := <-dataChanSelect
			switch chanSelect {
			case 0:
				readDataTransfer(clientData, serverData0)
			case 1:
				readDataTransfer(clientData, serverData1)
			case 2:
				readDataTransfer(clientData, serverData2)
			default:
				readDataTransfer(clientData, serverData3)
			}
		}
	}()

	// Use independent read addresses for efficient implementation.
	var readAddr0 protocol.Addr
	var readAddr1 protocol.Addr
	var readAddr2 protocol.Addr
	var readAddr3 protocol.Addr
	for {
		select {
		case readAddr0 = <-serverAddr0:
			clientAddr <- readAddr0
			dataChanSelect <- 0
		case readAddr1 = <-serverAddr1:
			clientAddr <- readAddr1
			dataChanSelect <- 1
		case readAddr2 = <-serverAddr2:
			clientAddr <- readAddr2
			dataChanSelect <- 2
		case readAddr3 = <-serverAddr3:
			clientAddr <- readAddr3
			dataChanSelect <- 3
		}
	}
}
