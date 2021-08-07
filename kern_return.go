// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && gc
// +build darwin,gc

package sys

// itoa converts val to a decimal string.
func itoa(val int) string {
	if val < 0 {
		return "-" + uitoa(uint(-val))
	}
	return uitoa(uint(val))
}

// uitoa converts val to a decimal string.
func uitoa(val uint) string {
	if val == 0 { // avoid string allocation
		return "0"
	}
	var buf [20]byte // big enough for 64bit value base 10
	i := len(buf) - 1
	for val >= 10 {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		i--
		val = q
	}
	// val < 10
	buf[i] = byte('0' + val)
	return string(buf[i:])
}

// Error returns a string representation of the KernReturn.
func (e KernReturn) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		s := errors[e]
		if s != "" {
			return s
		}
	}
	return "errno " + itoa(int(e))
}

// KernErrno returns common boxed Errno values, to prevent
// allocations at runtime.
func KernErrno(e KernReturn) error {
	switch e {
	case 0:
		return KernSuccess
	default:
		return KernReturn(e)
	}
}

// KernReturn Error table.
var errors = [...]string{
	0x1:   "specified address is not currently valid",
	0x2:   "specified memory is valid, but does not permit therequired forms of access",
	0x3:   "address range specified is already in use, or no address range of the size specified could be found",
	0x4:   "function requested was not applicable to this type of argument, or an argument is invalid",
	0x5:   "function could not be performed. A catch-all",
	0x6:   "system resource could not be allocated to fulfill this request",
	0x7:   "bogus access restriction",
	0x8:   "task in question does not hold receive rights for the port argument",
	0x9:   "during a page fault, the target address refers to a memory object that has been destroyed",
	0xa:   "during a page fault, the memory object indicated that the data could not be returned",
	0xb:   "receive right is already a member of the portset",
	0xc:   "receive right is not a member of a port set",
	0xd:   "name already denotes a right in the task",
	0xe:   "operation was aborted",
	0xf:   "name doesn't denote a right in the task",
	0x10:  "target task isn't an active task",
	0x11:  "name denotes a right, but not an appropriate right",
	0x12:  "blatant range error",
	0x13:  "operation would overflow limit on user-references",
	0x14:  "supplied (port) capability is improper",
	0x15:  "task already has send or receive rights for the port under another name",
	0x16:  "target host isn't actually a host",
	0x18:  "attempt was made to supply precious data for memory that is already present in a memory object",
	0x19:  "page was requested of a memory manager via memory_object_data_request for an object using a MEMORY_OBJECT_COPY_CALL strategy",
	0x1a:  "strategic copy was attempted of an object upon which a quicker copy is now possible",
	0x1b:  "argument applied to assert processor set privilege was not a processor set control port",
	0x1c:  "specified scheduling attributes exceed the thread's limits",
	0x1d:  "specified scheduling policy is not currently enabled for the processor set",
	0x1e:  "external memory manager failed to initialize the memory object",
	0x1f:  "thread is attempting to wait for an event for which there is already a waiting thread",
	0x20:  "attempt was made to destroy the default processor set",
	0x21:  "attempt was made to fetch an exception port that is protected, or to abort a thread while processing a protected exception",
	0x22:  "ledger was required but not supplied",
	0x23:  "port was not a memory cache control port",
	0x24:  "argument supplied to assert security privilege was not a host security port",
	0x25:  "thread_depress_abort was called on a thread which was not currently depressed",
	0x26:  "object has been terminated and is no longer available",
	0x27:  "lock set has been destroyed and is no longer available",
	0x28:  "thread holding the lock terminated before releasing the lock",
	0x29:  "lock is already owned by another thread",
	0x2a:  "lock is already owned by the calling thread",
	0x2b:  "semaphore has been destroyed and is no longer available",
	0x2c:  "return from RPC indicating the target server was terminated before it successfully replied",
	0x2d:  "terminate an orphaned activation",
	0x2e:  "allow an orphaned activation to continue executing",
	0x2f:  "empty thread activation (No thread linked to it)",
	0x30:  "remote node down or inaccessible",
	0x31:  "signalled thread was not actually waiting",
	0x32:  "some thread-oriented operation (semaphore_wait) timed out",
	0x33:  "during a page fault, indicates that the page was rejected as a result of a signature check",
	0x34:  "requested property cannot be changed at this time",
	0x35:  "provided buffer is of insufficient size for the requested data",
	0x36:  "KC on which the function is operating is missing",
	0x37:  "KC on which the function is operating is invalid",
	0x100: "maximum return value allowable",
}
