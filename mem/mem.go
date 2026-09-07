// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package mem provides memory layout information for Go unikernels compiled
// with `GOOS=tamago` as supported by the TamaGo framework for bare metal Go,
// see https://github.com/usbarmory/tamago.
package mem

import "unsafe"

var (
	// RamStart defines the start address of the physical or virtual memory
	// available to the runtime for allocation (including the code segment
	// which must be mapped within).
	RamStart uintptr

	// RamSize defines the total size of the physical or virtual memory
	// available to the runtime for allocation (including the code segment
	// which must be mapped within).
	RamSize uintptr

	// RamStackOffset, defines the negative offset from the end of the
	// available memory for stack allocation.
	RamStackOffset uintptr
)

// Region returns the start and end addresses of the physical RAM assigned to
// the Go runtime.
func Region() (start uintptr, end uintptr) {
	return RamStart, RamStart + RamSize
}

//go:linkname text runtime.text
var text uintptr

//go:linkname etext runtime.etext
var etext uintptr

// Text() returns the start and end addresses of the physical RAM containing
// the Go runtime global symbols.
func Text() (start uintptr, end uintptr) {
	start = uintptr(unsafe.Pointer(&text))
	end = uintptr(unsafe.Pointer(&etext))
	return
}
