// AI Foundry ET-SoC-1 Minion initialization
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramstart

package minion

import (
	_ "unsafe"

	"github.com/usbarmory/tamago/soc/aifoundry/etsoc1"
)

//go:linkname ramStart github.com/usbarmory/tamago/goos.RamStart
var ramStart uintptr = etsoc1.DRAM_BASE
