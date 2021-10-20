// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys_test

import (
	"flag"
	"runtime/debug"

	"github.com/go-darwin/sys"
)

var flagQuick = flag.Bool("quick", false, "skip slow tests, for second run in all.bash")

func init() {
	// We're testing the runtime, so make tracebacks show things
	// in the runtime. This only raises the level, so it won't
	// override GOTRACEBACK=crash from the user.
	SetTracebackEnv("system")
}

var traceback_cache uint32 = 2 << sys.TracebackShift
var traceback_env uint32

// SetTracebackEnv is like runtime/debug.SetTraceback, but it raises
// the "environment" traceback level, so later calls to
// debug.SetTraceback (e.g., from testing timeouts) can't lower it.
func SetTracebackEnv(level string) {
	debug.SetTraceback(level)
	traceback_env = traceback_cache
}
