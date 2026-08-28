// Microchip LAN969x configuration and support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package lan969x

import (
	_ "unsafe"
)

//go:linkname ramStackOffset github.com/usbarmory/tamago/goospkg.RamStackOffset
var ramStackOffset uintptr = 0x100

// Init takes care of the lower level initialization triggered early in runtime
// setup (e.g. internal/runtime/goospkg.InitHW1).
func Init() {
	ARM64.Init()
	ARM64.EnableCache()

	initTimers()
}
