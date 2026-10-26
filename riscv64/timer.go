// RISC-V 64-bit processor support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package riscv64

// GetTime returns the real or system time in nanoseconds.
func (cpu *CPU) GetTime(monotonic bool) (ns int64) {
	ns = int64(float64(cpu.Counter()) * cpu.TimerMultiplier)

	if monotonic {
		return
	}

	return ns + cpu.TimerOffset
}

// SetTime adjusts the system time to the argument nanoseconds value.
func (cpu *CPU) SetTime(ns int64) {
	if cpu.TimerMultiplier == 0 {
		return
	}

	cpu.TimerOffset = ns - int64(float64(cpu.Counter())*cpu.TimerMultiplier)
}
