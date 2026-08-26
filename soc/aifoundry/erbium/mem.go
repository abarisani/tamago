// AI Foundry Erbium initialization
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramstart

package erbium

import (
	_ "unsafe"
)

//go:linkname ramStart github.com/usbarmory/tamago/goos.RamStart
var ramStart uintptr = 0x40000000
