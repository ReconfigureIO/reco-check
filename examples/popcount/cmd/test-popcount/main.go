package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"xcl"
)

func main() {
	world := xcl.NewWorld()
	defer world.Release()

	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	buff := world.Malloc(xcl.WriteOnly, 4)
	defer buff.Free()

	krnl.SetArg(0, 12)
	krnl.SetMemoryArg(1, buff)

	krnl.Run(1, 1, 1)

	var ret uint32
	err := binary.Read(buff.Reader(), binary.LittleEndian, &ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Printf("%d\n", ret)
	if ret != 2 {
		os.Exit(1)
	}
}
