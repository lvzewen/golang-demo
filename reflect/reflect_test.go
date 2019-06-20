package main

import "testing"

// BenchmarkContains ...
func BenchmarkContains(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
	    Contains(sa, "r")
	}
 }
 
// BenchmarkStringContains ...
func BenchmarkStringContains(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringContains(sa, "r")
	}
}
