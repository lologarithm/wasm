// Copyright 2016 The wasm Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm_test

import (
	"fmt"
	"testing"

	"github.com/sbinet/wasm"
)

func TestOpen(t *testing.T) {
	mod, err := wasm.Open("testdata/hello.wasm")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("module header: %v\n", mod.Header)
	fmt.Printf("#sections: %d\n", len(mod.Sections))
}
