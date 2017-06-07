# go-teak/sdaccel

Directory of FPGA side code to be bundled as part of the build process.

## Conventions

The `verilog` should contain all verilog source to be bundled (including the wrappers to setup wires, go signals, etc).

There should be a `main.go` that includes required verilog source.
