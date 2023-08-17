package main

import (
	_ "embed"
)

// try using the icns file?

//go:embed build/macicon.tiff
var iconMac []byte

//go:embed build/maciconinactive.tiff
var iconMacInactive []byte

//go:embed build/othersicon.png
var iconOther []byte

//go:embed build/othersiconinactive.png
var iconOtherInactive []byte
