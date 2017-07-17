// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build !windows

package libkb

// /* Bionic has res_init() but it's not in any header */
// #ifdef __BIONIC__
// int res_init (void);
// #else
// #include<resolv.h>
// #endif
import "C"

func resInit() {
	C.res_init()
}
