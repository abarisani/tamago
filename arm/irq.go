// ARM processor support
// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package arm

// defined in irq.s
func irq_enable()
func irq_disable()

// EnableInterrupts enables IRQ and FIQ interrupts.
func (cpu *CPU) EnableInterrupts() {
	irq_enable()
}

// DisableInterrupts disables IRQ and FIQ interrupts.
func (cpu *CPU) DisableInterrupts() {
	irq_disable()
}
