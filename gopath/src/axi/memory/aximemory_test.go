package memory

import (
	"github.com/ReconfigureIO/sdaccel/axi/protocol"
	"testing"
	"testing/quick"
)

func TestWriteBurstUInt32(t *testing.T) {
	memWriteAddr := make(chan protocol.Addr)
	memWriteData := make(chan protocol.WriteData)
	memWriteResp := make(chan protocol.WriteResp)

	exit := make(chan struct{})
	dataPackets := 0

	fake := func() {
		for {
			select {
			case <-exit:
				return
			case a := <-memWriteAddr:
				// still loop if we get a 0
				for i := 0; i <= int(a.Len); i++ {
					<-memWriteData
					dataPackets += 1
				}
				memWriteResp <- protocol.WriteResp{false, [2]bool{true, true}, false}
			}
		}
	}

	f := func(reps uint16) bool {
		dataPackets = 0
		go fake()

		data := make(chan uint32)
		go func() {
			for i := 0; i < int(reps); i++ {
				data <- 1
			}
		}()

		WriteBurstUInt32(
			memWriteAddr, memWriteData, memWriteResp, true, 0, uint32(reps), data)

		exit <- struct{}{}
		return dataPackets == int(reps)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
