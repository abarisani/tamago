// Firecracker microvm support for tamago/amd64
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !linkramsize

package microvm

import (
	_ "unsafe"
)

// Applications can override ramSize with the `linkramsize` build tag.
//
// This is useful when large DMA descriptors are required to re-initialize
// tamago `dma` package in external RAM.

//go:linkname ramSize github.com/usbarmory/tamago/mem.RamSize
var ramSize uintptr = 0xb0000000 - dmaSize // 2560 MiB
