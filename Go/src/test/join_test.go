package main

import (
	"strconv"
	"testing"
)

func BenchmarkAddJoin(b *testing.B) {
	s := make([]string, 100)
	for index := 0; index < 100; index++ {
		s[index] = "arg" + strconv.Itoa(index)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addJoinStr(s)
	}
}

func BenchmarkJoin(b *testing.B) {
	s := make([]string, 100)
	for index := 0; index < 100; index++ {
		s[index] = "arg" + strconv.Itoa(index)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinStr(s)
	}
}

func TestJoin(t *testing.T) {
	ret := joinStr([]string{"1", "2"})
	if ret != "1 2" {
		t.Error("xxxxxxxx")
	}
}
