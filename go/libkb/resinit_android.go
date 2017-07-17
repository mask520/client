// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build android

package libkb

// /* Bionic has res_init() but it's not in any header */
// int res_init (void);
import "C"

func resInit() {
	C.res_init()
}
