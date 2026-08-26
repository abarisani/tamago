// NXP i.MX6UL initialization
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramstart

package imx6ul

import (
	_ "unsafe"
)

//go:linkname ramStart github.com/usbarmory/tamago/goos.RamStart
var ramStart uintptr = MMDC_BASE
