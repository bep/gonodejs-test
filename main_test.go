package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestBabelHTTP(t *testing.T) {
	nodeCmd := startNode()
	defer func() {
		nodeCmd.Process.Signal(os.Kill)
	}()
	time.Sleep(250 * time.Millisecond)

	s, err := transpileViaHTTP()
	if err != nil {
		t.Error(err)
	}

	// Sanity check
	if !strings.Contains(s, "return n + 1;") {
		t.Errorf("got: %s", s)
	}

}

func TestBabelExec(t *testing.T) {

	s, err := transpileViaExec()
	if err != nil {
		t.Error(err)
	}

	// Sanity check
	if !strings.Contains(s, "return n + 1;") {
		t.Errorf("got: %s", s)
	}

}

func BenchmarkBabelHTTP(b *testing.B) {
	nodeCmd := startNode()
	defer func() {
		nodeCmd.Process.Signal(os.Kill)
	}()

	// We could probably improve this.
	time.Sleep(250 * time.Millisecond)

	// Ignore the server startup in the timings. Not entirely realistic,
	// but it matches the intended use as we can start the server while
	// doing other stuff.
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		s, err := transpileViaHTTP()
		if err != nil {
			b.Fatal(err)
		}
		if !strings.Contains(s, "return n + 1;") {
			b.Fatalf("got: %s", s)
		}
	}
}

func BenchmarkBabelExec(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s, err := transpileViaExec()
		if err != nil {
			b.Fatal(err)
		}
		if !strings.Contains(s, "return n + 1;") {
			b.Fatalf("got: %s", s)
		}
	}
}
