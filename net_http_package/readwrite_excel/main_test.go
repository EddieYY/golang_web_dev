package main

import "testing"

var excelFileName = "/Users/yychenaa//Documents/go/src/github.com/EddieYY/test.xlsx"

/*
func TestReadxlsxChan(t *testing.T) {
	readxlsxChan(excelFileName)
}
*/
func BenchmarkReadxlsxChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readxlsxChan(excelFileName)
	}

}

func BenchmarkReadxlsx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readxlsx(excelFileName)
	}

}
func BenchmarkReadxlsxslice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readxlsxslice(excelFileName)
	}

}
func BenchmarkReadxlsxChanSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readxlsxChanSlice(excelFileName)
	}

}
