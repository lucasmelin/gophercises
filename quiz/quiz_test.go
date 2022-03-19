package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadCsv(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("5+5,10")
	content, err := readCsvFile(&buffer)
	if err != nil {
		t.Error("Failed to read csv data")
	}
	want := make([][]string, 0)
	want = append(want, []string{"5+5", "10"})
	if !reflect.DeepEqual(want, content) {
		t.Errorf("Got %v, want %v", content, want)
	}

}
