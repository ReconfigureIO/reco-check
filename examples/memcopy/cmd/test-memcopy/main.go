package main

import (
	"bytes"
	"encoding/binary"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"log"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

const DATA_WIDTH = 12

func main() {
	var conf = quick.Config{Rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano())), MaxCount: 1}

	world := xcl.NewWorld()
	defer world.Release()

	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	memcpy := func(input [DATA_WIDTH]uint64) bool {

		byteLength := uint(binary.Size(input))

		outputBuff := world.Malloc(xcl.WriteOnly, byteLength)
		defer outputBuff.Free()

		inputBuff := world.Malloc(xcl.ReadOnly, byteLength)
		defer inputBuff.Free()

		inToBytes := new(bytes.Buffer)
		binary.Write(inToBytes, binary.LittleEndian, &input)
		inputBuff.Write(inToBytes.Bytes())

		krnl.SetMemoryArg(0, inputBuff)
		krnl.SetMemoryArg(1, outputBuff)
		krnl.SetArg(2, uint32(len(input)))

		krnl.Run(1, 1, 1)

		resp := make([]byte, byteLength)
		outputBuff.Read(resp)

		var ret [DATA_WIDTH]uint64
		err := binary.Read(bytes.NewReader(resp), binary.LittleEndian, &ret)

		if err != nil {
			log.Fatal("binary.Read failed:", err)
		}

		if !reflect.DeepEqual(ret, input) {
			log.Printf("%v != %v", ret, input)
			return false
		}
		return true
	}

	if err := quick.Check(memcpy, &conf); err != nil {
		log.Fatal(err)
	}
}
