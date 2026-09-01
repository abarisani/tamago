// Linux user space support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package linux_user provides support for using `GOOS=tamago` in Linux user
// space.
//
// This package is only meant to be used with `GOOS=tamago` as supported by the
// TamaGo framework for bare metal Go, see https://github.com/usbarmory/tamago.
package linux_user

import (
	_ "unsafe"

	"github.com/usbarmory/tamago/goos"
)

const (
	ramStart       uint = 0x80000000
	ramSize        uint = 0x20000000 // 512MB
	ramStackOffset uint = 0x100
)

// defined in syscall_*.s
func sys_exit(code int32)
func sys_write(c *byte)
func sys_clock_gettime() (ns int64)
func sys_getrandom(b []byte, n int)

//go:linkname nanotime internal/runtime/goospkg.Nanotime
func nanotime() int64 {
	return sys_clock_gettime()
}

//go:linkname initRNG internal/runtime/goospkg.InitRNG
func initRNG() {}

//go:linkname getRandomData internal/runtime/goospkg.GetRandomData
func getRandomData(b []byte) {
	sys_getrandom(b, len(b))
}

// preallocated memory to avoid malloc during panic
var a [1]byte

//go:linkname printk internal/runtime/goospkg.WriteConsole
func printk(c byte) {
	a[0] = c
	sys_write(&a[0])
}

//go:linkname inithw0 internal/runtime/goospkg.InitHW0
func inithw0() {
	goos.Bloc = uintptr(ramStart)
}

//go:linkname inithw1 internal/runtime/goospkg.InitHW1
func inithw1() {
	goos.Exit = sys_exit
}
