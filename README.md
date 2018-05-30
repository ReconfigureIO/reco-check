# This project has moved; this repository is no longer maintained.

It is no longer maintained here, it is now maintained as part of our internal compiler infrastructure.

---

# reco-check

A program to validate Go files for compatibility with Reconfigure.io's Go compiler.

## Installing

Make sure that all files located in `bins` are located on your `%PATH%`.
Make sure that your `GOPATH` is seto to the `gopath` folder.

```
set PATH=%PATH%;%CD%\bins
set GOPATH=%CD%\gopath
```

## Usage

`reco-check main.go`

## Examples

This repository includes several examples, located in `example`. You
should be able to use these as starting points for your own
designs. They can be checked like so:

`reco-check examples/addition/main.go`

## Further Reading

See our [online documentation](http://docs.reconfigure.io/).

The tutorials described will follow roughly the same guideline, but
you can use the examples included here. Where you would do `reco
build`, you can instead run `reco-check main.go`.

## Licensing

`bins/reco-check.exe` is available through the Permissive Binary
License. A copy is provided in LICENSE.

`bins/goblin.exe` is a prebuilt version of
[goblin](https://github.com/ReconfigureIO/goblin), and is availabe
under its license.

All source code is availabe under the Apache License Version 2.0. A
copy is included at SOURCE_LICENSE
