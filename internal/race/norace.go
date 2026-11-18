// Copyright (c) 2026 malasahjagotwin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !race

package race

import (
	"unsafe"
)

const Enabled = false

func Acquire(addr unsafe.Pointer) {
}

func Release(addr unsafe.Pointer) {
}

func ReleaseMerge(addr unsafe.Pointer) {
}

func Disable() {
}

func Enable() {
}

func Read(addr unsafe.Pointer) {
}

func Write(addr unsafe.Pointer) {
}

func ReadRange(addr unsafe.Pointer, len int) {
}

func WriteRange(addr unsafe.Pointer, len int) {
}

func Errors() int { return 0 }
