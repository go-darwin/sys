// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	_ "runtime" // for go:linkname
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
	"unicode/utf8"
	_ "unsafe" // for go:linkname

	"github.com/go-darwin/sys/testenv"
)

// Strings and slices that don't escape and fit into tmpBuf are stack allocated,
// which defeats using AllocsPerRun to test other optimizations.
const sizeNoStack = 100

func BenchmarkCompareStringEqual(b *testing.B) {
	bytes := []byte("Hello Gophers!")
	s1, s2 := string(bytes), string(bytes)
	for i := 0; i < b.N; i++ {
		if s1 != s2 {
			b.Fatal("s1 != s2")
		}
	}
}

func BenchmarkCompareStringIdentical(b *testing.B) {
	s1 := "Hello Gophers!"
	s2 := s1
	for i := 0; i < b.N; i++ {
		if s1 != s2 {
			b.Fatal("s1 != s2")
		}
	}
}

func BenchmarkCompareStringSameLength(b *testing.B) {
	s1 := "Hello Gophers!"
	s2 := "Hello, Gophers"
	for i := 0; i < b.N; i++ {
		if s1 == s2 {
			b.Fatal("s1 == s2")
		}
	}
}

func BenchmarkCompareStringDifferentLength(b *testing.B) {
	s1 := "Hello Gophers!"
	s2 := "Hello, Gophers!"
	for i := 0; i < b.N; i++ {
		if s1 == s2 {
			b.Fatal("s1 == s2")
		}
	}
}

func BenchmarkCompareStringBigUnaligned(b *testing.B) {
	bytes := make([]byte, 0, 1<<20)
	for len(bytes) < 1<<20 {
		bytes = append(bytes, "Hello Gophers!"...)
	}
	s1, s2 := string(bytes), "hello"+string(bytes)
	for i := 0; i < b.N; i++ {
		if s1 != s2[len("hello"):] {
			b.Fatal("s1 != s2")
		}
	}
	b.SetBytes(int64(len(s1)))
}

func BenchmarkCompareStringBig(b *testing.B) {
	bytes := make([]byte, 0, 1<<20)
	for len(bytes) < 1<<20 {
		bytes = append(bytes, "Hello Gophers!"...)
	}
	s1, s2 := string(bytes), string(bytes)
	for i := 0; i < b.N; i++ {
		if s1 != s2 {
			b.Fatal("s1 != s2")
		}
	}
	b.SetBytes(int64(len(s1)))
}

func BenchmarkConcatStringAndBytes(b *testing.B) {
	s1 := []byte("Gophers!")
	for i := 0; i < b.N; i++ {
		_ = "Hello " + string(s1)
	}
}

var escapeString string

func BenchmarkSliceByteToString(b *testing.B) {
	buf := []byte{'!'}
	for n := 0; n < 8; n++ {
		b.Run(strconv.Itoa(len(buf)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				escapeString = string(buf)
			}
		})
		buf = append(buf, buf...)
	}
}

var stringdata = []struct{ name, data string }{
	{"ASCII", "01234567890"},
	{"Japanese", "日本語日本語日本語"},
	{"MixedLength", "$Ѐࠀက퀀𐀀\U00040000\U0010FFFF"},
}

var sinkInt int

func BenchmarkRuneCount(b *testing.B) {
	// Each sub-benchmark counts the runes in a string in a different way.
	b.Run("lenruneslice", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sinkInt += len([]rune(sd.data))
				}
			})
		}
	})
	b.Run("rangeloop", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					n := 0
					for range sd.data {
						n++
					}
					sinkInt += n
				}
			})
		}
	})
	b.Run("utf8.RuneCountInString", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sinkInt += utf8.RuneCountInString(sd.data)
				}
			})
		}
	})
}

func BenchmarkRuneIterate(b *testing.B) {
	b.Run("range", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for range sd.data {
					}
				}
			})
		}
	})
	b.Run("range1", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for range sd.data {
					}
				}
			})
		}
	})
	b.Run("range2", func(b *testing.B) {
		for _, sd := range stringdata {
			b.Run(sd.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for range sd.data {
					}
				}
			})
		}
	})
}

func BenchmarkArrayEqual(b *testing.B) {
	a1 := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	a2 := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if a1 != a2 {
			b.Fatal("not equal")
		}
	}
}

var toRemove []string

func TestMain(m *testing.M) {
	status := m.Run()
	for _, file := range toRemove {
		os.RemoveAll(file)
	}
	os.Exit(status)
}

var testprog struct {
	sync.Mutex
	dir    string
	target map[string]buildexe
}

type buildexe struct {
	exe string
	err error
}

func buildTestProg(t *testing.T, binary string, flags ...string) (string, error) {
	if *flagQuick {
		t.Skip("-quick")
	}

	testprog.Lock()
	defer testprog.Unlock()
	if testprog.dir == "" {
		dir, err := os.MkdirTemp("", "go-build")
		if err != nil {
			t.Fatalf("failed to create temp directory: %v", err)
		}
		testprog.dir = dir
		toRemove = append(toRemove, dir)
	}

	if testprog.target == nil {
		testprog.target = make(map[string]buildexe)
	}
	name := binary
	if len(flags) > 0 {
		name += "_" + strings.Join(flags, "_")
	}
	target, ok := testprog.target[name]
	if ok {
		return target.exe, target.err
	}

	exe := filepath.Join(testprog.dir, name+".exe")
	cmd := exec.Command(testenv.GoToolPath(t), append([]string{"build", "-o", exe}, flags...)...)
	cmd.Dir = "testdata/" + binary
	out, err := testenv.CleanCmdEnv(cmd).CombinedOutput()
	if err != nil {
		target.err = fmt.Errorf("building %s %v: %v\n%s", binary, flags, err, out)
		testprog.target[name] = target
		return "", target.err
	}
	target.exe = exe
	testprog.target[name] = target
	return exe, nil
}

func runTestProg(t *testing.T, binary, name string, env ...string) string {
	if *flagQuick {
		t.Skip("-quick")
	}

	testenv.MustHaveGoBuild(t)

	exe, err := buildTestProg(t, binary)
	if err != nil {
		t.Fatal(err)
	}

	return runBuiltTestProg(t, exe, name, env...)
}

// sigquit is the signal to send to kill a hanging testdata program.
// Send SIGQUIT to get a stack trace.
var sigquit = syscall.SIGQUIT

func runBuiltTestProg(t *testing.T, exe, name string, env ...string) string {
	if *flagQuick {
		t.Skip("-quick")
	}

	testenv.MustHaveGoBuild(t)

	cmd := testenv.CleanCmdEnv(exec.Command(exe, name))
	cmd.Env = append(cmd.Env, env...)
	if testing.Short() {
		cmd.Env = append(cmd.Env, "RUNTIME_TEST_SHORT=1")
	}
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Start(); err != nil {
		t.Fatalf("starting %s %s: %v", exe, name, err)
	}

	// If the process doesn't complete within 1 minute,
	// assume it is hanging and kill it to get a stack trace.
	p := cmd.Process
	done := make(chan bool)
	go func() {
		scale := 1
		// This GOARCH/GOOS test is copied from cmd/dist/test.go.
		// TODO(iant): Have cmd/dist update the environment variable.
		if runtime.GOARCH == "arm" || runtime.GOOS == "windows" {
			scale = 2
		}
		if s := os.Getenv("GO_TEST_TIMEOUT_SCALE"); s != "" {
			if sc, err := strconv.Atoi(s); err == nil {
				scale = sc
			}
		}

		select {
		case <-done:
		case <-time.After(time.Duration(scale) * time.Minute):
			p.Signal(sigquit)
		}
	}()

	if err := cmd.Wait(); err != nil {
		t.Logf("%s %s exit status: %v", exe, name, err)
	}
	close(done)

	return b.String()
}

func TestLargeStringConcat(t *testing.T) {
	output := runTestProg(t, "testprog", "stringconcat")
	want := "panic: " + strings.Repeat("0", 1<<10) + strings.Repeat("1", 1<<10) +
		strings.Repeat("2", 1<<10) + strings.Repeat("3", 1<<10)
	if !strings.HasPrefix(output, want) {
		t.Fatalf("output does not start with %q:\n%s", want, output)
	}
}

func TestCompareTempString(t *testing.T) {
	s := strings.Repeat("x", sizeNoStack)
	b := []byte(s)
	n := testing.AllocsPerRun(1000, func() {
		if string(b) != s {
			t.Fatalf("strings are not equal: '%v' and '%v'", string(b), s)
		}
		if string(b) == s {
		} else {
			t.Fatalf("strings are not equal: '%v' and '%v'", string(b), s)
		}
	})
	if n != 0 {
		t.Fatalf("want 0 allocs, got %v", n)
	}
}

func TestStringIndexHaystack(t *testing.T) {
	// See issue 25864.
	haystack := []byte("hello")
	needle := "ll"
	n := testing.AllocsPerRun(1000, func() {
		if strings.Index(string(haystack), needle) != 2 {
			t.Fatalf("needle not found")
		}
	})
	if n != 0 {
		t.Fatalf("want 0 allocs, got %v", n)
	}
}

func TestStringIndexNeedle(t *testing.T) {
	// See issue 25864.
	haystack := "hello"
	needle := []byte("ll")
	n := testing.AllocsPerRun(1000, func() {
		if strings.Index(haystack, string(needle)) != 2 {
			t.Fatalf("needle not found")
		}
	})
	if n != 0 {
		t.Fatalf("want 0 allocs, got %v", n)
	}
}

func TestStringOnStack(t *testing.T) {
	s := ""
	for i := 0; i < 3; i++ {
		s = "a" + s + "b" + s + "c"
	}

	if want := "aaabcbabccbaabcbabccc"; s != want {
		t.Fatalf("want: '%v', got '%v'", want, s)
	}
}

func TestIntString(t *testing.T) {
	// Non-escaping result of intstring.
	s := ""
	for i := rune(0); i < 4; i++ {
		s += string(i+'0') + string(i+'0'+1)
	}
	if want := "01122334"; s != want {
		t.Fatalf("want '%v', got '%v'", want, s)
	}

	// Escaping result of intstring.
	var a [4]string
	for i := rune(0); i < 4; i++ {
		a[i] = string(i + '0')
	}
	s = a[0] + a[1] + a[2] + a[3]
	if want := "0123"; s != want {
		t.Fatalf("want '%v', got '%v'", want, s)
	}
}

func TestIntStringAllocs(t *testing.T) {
	unknown := '0'
	n := testing.AllocsPerRun(1000, func() {
		s1 := string(unknown)
		s2 := string(unknown + 1)
		if s1 == s2 {
			t.Fatalf("bad")
		}
	})
	if n != 0 {
		t.Fatalf("want 0 allocs, got %v", n)
	}
}

func TestRangeStringCast(t *testing.T) {
	s := strings.Repeat("x", sizeNoStack)
	n := testing.AllocsPerRun(1000, func() {
		for i, c := range []byte(s) {
			if c != s[i] {
				t.Fatalf("want '%c' at pos %v, got '%c'", s[i], i, c)
			}
		}
	})
	if n != 0 {
		t.Fatalf("want 0 allocs, got %v", n)
	}
}

func isZeroed(b []byte) bool {
	for _, x := range b {
		if x != 0 {
			return false
		}
	}
	return true
}

func isZeroedR(r []rune) bool {
	for _, x := range r {
		if x != 0 {
			return false
		}
	}
	return true
}

func TestString2Slice(t *testing.T) {
	// Make sure we don't return slices that expose
	// an unzeroed section of stack-allocated temp buf
	// between len and cap. See issue 14232.
	s := "foož"
	b := ([]byte)(s)
	if !isZeroed(b[len(b):cap(b)]) {
		t.Errorf("extra bytes not zeroed")
	}
	r := ([]rune)(s)
	if !isZeroedR(r[len(r):cap(r)]) {
		t.Errorf("extra runes not zeroed")
	}
}
