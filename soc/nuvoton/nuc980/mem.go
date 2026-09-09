// Nuvoton NUC980 SoC support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramstart

package nuc980

import (
	_ "unsafe"
)

//go:linkname ramStart github.com/usbarmory/tamago/mem.RamStart
var ramStart uintptr = SDRAM_BASE

//go:linkname textStart internal/runtime/goospkg.TextAddr
var textStart uintptr = SDRAM_BASE + 0x10000
