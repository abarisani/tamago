// Custom GOOS support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build !tiny && !shared_stack

package goos

// Required constants.
const (
	// StackSystem is a number of additional bytes to add to each stack
	// below the usual guard area.
	StackSystem = 512
)
