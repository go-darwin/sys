// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

package sys

import "syscall"

// Errno is the type alias of syscall.Errno, an unsigned number describing an error condition.
//
// It implements the error interface. The zero Errno is by convention
// a non-error, so code to convert from Errno to error should use:
//
//	err = nil
//	if errno != 0 {
//		err = errno
//	}
//
// Errno values can be tested against error values from the os package
// using errors.Is. For example:
//
//	_, _, err := syscall.Syscall(...)
//	if errors.Is(err, fs.ErrNotExist) ...
type Errno = syscall.Errno
