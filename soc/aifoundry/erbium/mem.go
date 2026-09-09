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

//go:linkname ramStart github.com/usbarmory/tamago/mem.RamStart
var ramStart uintptr = MRAM_BASE

//go:linkname textStart internal/runtime/goospkg.TextAddr
var textStart uintptr = MRAM_BASE + 0x10000
