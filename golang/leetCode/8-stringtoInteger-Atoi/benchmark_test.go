package main

import "testing"

func BenchmarkReverse(b *testing.B) {
	input := "18446744073709551617"
	for i := 0; i < b.N; i++ {
		myAtoi(input)
	}
}

func BenchmarkV1MyAtoi(b *testing.B) {
	input := "18446744073709551617"
	for i := 0; i < b.N; i++ {
		v1MyAtoi(input)
	}
}
