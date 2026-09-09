// BCM2835 SoC support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramstart

package bcm2835

import (
	_ "unsafe"
)

//go:linkname ramStart github.com/usbarmory/tamago/mem.RamStart
var ramStart uintptr = 0x00100000

//go:linkname textStart internal/runtime/goospkg.TextAddr
var textStart uintptr = 0x00110000
