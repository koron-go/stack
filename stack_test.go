package stack_test

import (
	"reflect"
	"testing"

	"github.com/koron-go/stack"
)

func stack0() []*stack.Frame {
	return stack.Frames(0)
}

func stackN(n int) []*stack.Frame {
	if n == 0 {
		return stack0()
	}
	return stackN(n - 1)
}

func testFrames(t *testing.T, n int) {
	ff0 := stackN(1)
	ff1 := stackN(n + 1)
	assertFrames(t, n, ff0, ff1)
}

func assertFrames(t *testing.T, n int, ff0, ff1 []*stack.Frame) {
	lead0 := ff0[0:3]
	lead1 := ff1[0:3]
	if !reflect.DeepEqual(lead0, lead1) {
		t.Fatalf("lead frames mismatch at #%d:\nff0: %s\nff1: %s\n", n,
			stack.Join(lead0, "\n\t"),
			stack.Join(lead1, "\n\t"))
	}

	m0 := len(ff0)
	m1 := len(ff1)
	trail0 := ff0[m0-3 : m0]
	trail1 := ff1[m0+n-3 : m1]
	if !reflect.DeepEqual(trail0, trail1) {
		t.Fatalf("trail frames mismatch at #%d:\nff0: %s\nff1: %s\n", n,
			stack.Join(trail0, "\n\t"),
			stack.Join(trail1, "\n\t"))
	}

	if n == 0 {
		return
	}

	mid0 := make([]*stack.Frame, n)
	for i := range mid0 {
		mid0[i] = lead0[2]
	}
	mid1 := ff1[3 : n+3]
	if !reflect.DeepEqual(mid0, mid1) {
		t.Fatalf("middle frames mismatch at #%d:\nff0: %s\nff1: %s\n", n,
			stack.Join(mid0, "\n\t"),
			stack.Join(mid1, "\n\t"))
	}
}

func TestFrames(t *testing.T) {
	for i := 0; i < 30; i++ {
		testFrames(t, i)
	}
}
