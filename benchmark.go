package main

import "testing"

func BenchmarkIface(b *testing.B) {
    for i := 0; i < b.N; i++ {
        logIface(debug, "test iface")
    }
}

func BenchmarkString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        logString(debug, "test string")
    }
}