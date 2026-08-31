// SiFive FU540 initialization
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package fu540

import (
	_ "unsafe"

	"github.com/usbarmory/tamago/riscv64"
)

//go:linkname ramStackOffset github.com/usbarmory/tamago/mem.RamStackOffset
var ramStackOffset uintptr = 0x100

// Init takes care of the lower level initialization triggered early in runtime
// setup (e.g. internal/runtime/goospkg.InitHW1).
func Init() {
	RV64.Init()

	riscv64.IPI = CLINT.IPI
	riscv64.ClearIPI = CLINT.ClearIPI
}
